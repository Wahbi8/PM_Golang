package rabbitmq

import (
	"encoding/json"

	"github.com/Wahbi8/PM_Golang/DTO"
	"github.com/Wahbi8/PM_Golang/Services"
	"github.com/Wahbi8/PM_Golang/logger"
	"github.com/Wahbi8/PM_Golang/repository"
	amqp "github.com/rabbitmq/amqp091-go"
)

func SendQueueEmail(emailInfo dto.EmailInfo) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnError(err, "Failed to create connection in sendEmail")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to create connection in send email")
	defer ch.Close()

	err = ch.Qos(5, 0, false)
	failOnError(err, "Failed to set QoS")

	q, err := ch.QueueDeclare(
		"email_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Faile to declare queue on SendEmail")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack (set to false for manual ack)
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register consumer on SendEmail")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var msg dto.EmailInfo
			err := json.Unmarshal(d.Body, &msg)
			if err != nil {
				logger.Log.Error().Err(err).Msg("Malformed message")
				d.Ack(false) // Remove bad message from queue
				continue
			}

			err = Services.SendEmail(msg.Recipient, msg.Subject, msg.Message)

			if err != nil {
				if msg.Retry >= 3 {
					logger.Log.Error().Str("recipient", msg.Recipient).Msg("Max retries reached, saving to DB")
					repository.InsertFailedMsgs(&msg, err.Error())
					d.Ack(false)
				} else {
					msg.Retry++
					logger.Log.Warn().Int("retry", msg.Retry).Str("recipient", msg.Recipient).Msg("Retrying email")

					// 1. Convert updated struct back to bytes
					newBody, _ := json.Marshal(msg)

					// 2. Publish the UPDATED message back to the queue
					err = ch.Publish("", q.Name, false, false, amqp.Publishing{
						ContentType:  "application/json",
						DeliveryMode: amqp.Persistent,
						Body:         newBody,
					})

					d.Ack(false)
				}
			} else {
				logger.Log.Info().Str("recipient", msg.Recipient).Msg("Email sent successfully")
				d.Ack(false)
			}
		}
	}()

	logger.Log.Info().Msg("Worker started. Waiting for messages...")
	<-forever
}
