package rabbitmq

import (
	// "context"
	"log"
	// "math/bits"
	// "os"
	// "strings"
	// "time"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/Wahbi8/PM_Golang/services"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panic("%s: %s", msg, err)
	}
}

func SendQueueMsg() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnError(err, "Failed to create connection")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare()
}

func QueueMsg(emailInfo Services.EmailInfo) []byte {
	message := map[string]interface{}{
			"to":      emailInfo.Recipient, 
			"subject": "Invoice",
			"body":    emailInfo.Message,
			"retry":   0,
		}

	bytes, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error encoding JSON: %s", err)
		return nil
	}

	return bytes
}