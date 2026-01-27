package rabbitmq

import (
	
	"encoding/json"
	"fmt"

	"github.com/Wahbi8/PM_Golang/Services"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/Wahbi8/PM_Golang/DTO"
	"github.com/Wahbi8/PM_Golang/repository"
)

func ResendFailedMsgs() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnError(err, "Failed to create connection in resend")
	defer conn.Close()
	
	ch, err := conn.Channel()
	failOnError(err, "Failed to create channel")
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
	failOnError(err, "Faile to declare queue on resendEmail")

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
		for {
			//TODO: i will recieve list of failed msg
			dbMsg := repository.GetFailedEmailsFromDB()
			
			body2, err := json.Marshal(dbMsg)
			if err != nil {
				fmt.Printf("Malformed message in db: %v\n", err)
				continue
			} else {
				err = ch.Publish("", q.Name, false, false, amqp.Publishing{
								ContentType:  "application/json",
								DeliveryMode: amqp.Persistent,
								Body:         body2,
							})
			}
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
					if msg.Retry > 3 {
						fmt.Printf("Failed: Max retries are reached for %s in resend \n", msg.Recipient)
						//TODO: create a db table to store the failed email to be processed manualy
						d.Ack(false)
					} else {
						msg.Retry++
						fmt.Printf("RETRYING (%d/3): Sending back to queue\n", msg.Retry)

						newBody, _ := json.Marshal(msg)
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
		}
	}()
	fmt.Println("Worker started. Waiting for messages...")
	<-forever
}