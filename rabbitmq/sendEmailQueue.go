package rabbitmq

import (
	"encoding/json"
	"fmt"

	"github.com/Wahbi8/PM_Golang/Services"
	amqp "github.com/rabbitmq/amqp091-go"
)

func SendEmail(emailInfo Services.EmailInfo) {
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
			var emailMsg map[string]string
			err := json.Unmarshal(d.Body, &emailMsg)
			if err != nil {
				fmt.Printf("Error parsing message: %v\n", err)
				d.Nack(false, false) // reject message
				continue
			}

			err = Services.SendEmail(emailMsg["to"], emailMsg["subject"], emailMsg["body"])
			if err != nil {
				fmt.Printf("Failed to send email: %v\n", err)
				d.Nack(false, true) // requeue message
			} else {
				fmt.Printf("Email sent to: %s\n", emailMsg["to"])
				d.Ack(false) // acknowledge message
			}
		}
	}()
	fmt.Println("Waiting for messages...")
	<-forever
}