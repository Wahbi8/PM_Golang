package rabbitmq

import (
	"encoding/json"

	"github.com/Wahbi8/PM_Golang/DTO"
	"github.com/Wahbi8/PM_Golang/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		logger.Log.Fatal().Err(err).Str("message", msg).Msg("Failed")
	}
}

func SendQueueMsg(emailInfo dto.EmailInfo) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnError(err, "Failed to create connection")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"email_queue",
		true,  //durable
		false, //delete when unused
		false, //exclusive
		false, //no-wait
		nil,   //arguments
	)
	failOnError(err, "problem with the queue")

	err = ch.Publish(
		"",     //exchange
		q.Name, //routing key (queue name)
		false,  //mandatory
		false,  //emmediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         QueueMsg(emailInfo),
		},
	)
	failOnError(err, "Failed to publish message")

	logger.Log.Info().Str("recipient", emailInfo.Recipient).Msg("Email queued")
}

func QueueMsg(emailInfo dto.EmailInfo) []byte {
	message := map[string]interface{}{
		"RecipientEmail": emailInfo.Recipient,
		"subject":        "Invoice",
		"Body":           emailInfo.Message,
		"retry":          emailInfo.Retry,
		"InvoiceId":      emailInfo.InvoiceId,
		"InvoiceType":    emailInfo.InvoiceType,
	}

	bytes, err := json.Marshal(message)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Error encoding JSON")
		return nil
	}

	return bytes
}
