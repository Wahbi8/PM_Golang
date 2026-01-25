package rabbitmq

import (
	"encoding/json"
	"fmt"

	"github.com/Wahbi8/PM_Golang/Services"
	amqp "github.com/rabbitmq/amqp091-go"
    "github.com/Wahbi8/PM_Golang/DTO"

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
				fmt.Printf("Malformed message: %v\n", err)
				d.Ack(false) // Remove bad message from queue
				continue
			}

			err = Services.SendEmail(msg.Recipient, msg.Subject, msg.Message)
			
			if err != nil {
				if msg.Retry >= 3 {
					fmt.Printf("FAILED: Max retries reached for %s. Saving to DB...\n", msg.Recipient)
					// TODO: Add to database here



					d.Ack(false)
				} else {
					msg.Retry++
					fmt.Printf("RETRYING (%d/3): Sending back to queue\n", msg.Retry)
					
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
				fmt.Printf("SUCCESS: Email sent to %s\n", msg.Recipient)
				d.Ack(false)
			}
		}
	}()

	fmt.Println("Worker started. Waiting for messages...")
	<-forever
}