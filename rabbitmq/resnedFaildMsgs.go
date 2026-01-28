package rabbitmq

import (
	
	"encoding/json"
	"fmt"
	"time"

	// "github.com/Wahbi8/PM_Golang/Services"
	amqp "github.com/rabbitmq/amqp091-go"
	// "github.com/Wahbi8/PM_Golang/DTO"
	"github.com/Wahbi8/PM_Golang/repository"
)

func ResendFailedMsgs() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnError(err, "Failed to create connection")
	defer conn.Close()
	
	ch, err := conn.Channel()
	failOnError(err, "Failed to create channel")
	defer ch.Close()

	ticker := time.NewTicker(60 * time.Second)

	fmt.Println("Resend Cron Job started...")

	for range ticker.C {
		q, err := ch.QueueDeclare(
			"email_queue",
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			fmt.Printf("Error inspecting queue: %v\n", err)
			continue
		}

		if q.Messages >= 5 {
			fmt.Printf("Queue already has %d messages. Skipping DB fetch to prevent overload.\n", q.Messages)
			continue
		}

		failedEmails, err := repository.GetFailedEmailsFromDB()
		if err != nil {
			fmt.Printf("DB Error: %v\n", err)
			continue
		}

		if len(failedEmails) == 0 {
			continue
		}

		for _, msg := range failedEmails {
			body, err := json.Marshal(msg)
			if err != nil {
				fmt.Printf("Failed to marshal msg: %v\n", err)
				continue
			}

			err = ch.Publish("", q.Name, false, false, amqp.Publishing{
				ContentType:  "application/json",
				DeliveryMode: amqp.Persistent,
				Body:         body,
			})

			if err != nil {
				fmt.Printf("Failed to publish: %v\n", err)
			} else {
				fmt.Printf("Successfully re-queued email for: %s\n", msg.Recipient)
				
				//TODO: Add function delete from db
				repository.DeleteFailedEmail(msg.InvoiceId) 
			}
		}
	}
}