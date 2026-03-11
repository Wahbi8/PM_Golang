package rabbitmq

import (
	"encoding/json"
	"time"

	"github.com/Wahbi8/PM_Golang/logger"
	"github.com/Wahbi8/PM_Golang/repository"
	amqp "github.com/rabbitmq/amqp091-go"
)

func ResendFailedMsgs() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnError(err, "Failed to create connection")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to create channel")
	defer ch.Close()

	ticker := time.NewTicker(60 * time.Second)

	logger.Log.Info().Msg("Resend cron job started")

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
			logger.Log.Error().Err(err).Msg("Error inspecting queue")
			continue
		}

		if q.Messages >= 5 {
			logger.Log.Warn().Int("queue_size", q.Messages).Msg("Queue overloaded, skipping DB fetch")
			continue
		}

		failedEmails, err := repository.GetFailedEmailsFromDB()
		if err != nil {
			logger.Log.Error().Err(err).Msg("DB error")
			continue
		}

		if len(failedEmails) == 0 {
			continue
		}

		for _, msg := range failedEmails {
			body, err := json.Marshal(msg)
			if err != nil {
				logger.Log.Error().Err(err).Msg("Failed to marshal message")
				continue
			}

			err = ch.Publish("", q.Name, false, false, amqp.Publishing{
				ContentType:  "application/json",
				DeliveryMode: amqp.Persistent,
				Body:         body,
			})

			if err != nil {
				logger.Log.Error().Err(err).Str("recipient", msg.Recipient).Msg("Failed to publish")
			} else {
				logger.Log.Info().Str("recipient", msg.Recipient).Msg("Re-queued failed email")
				repository.DeleteFailedEmail(msg.InvoiceId)
			}
		}
	}
}
