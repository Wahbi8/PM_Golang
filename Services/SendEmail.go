package Services

import (
	"github.com/Wahbi8/PM_Golang/logger"
	"github.com/joho/godotenv"
	"github.com/resend/resend-go/v2"
	"os"
)

func SendEmail(to, subject, body string) error {
	godotenv.Load()

	apiKey := os.Getenv("Resend_api_key")
	if apiKey == "" {
		logger.Log.Fatal().Msg("Resend API key not found")
	}

	client := resend.NewClient(apiKey)

	params := &resend.SendEmailRequest{
		From:    "Acme <onboarding@resend.dev>",
		To:      []string{"wahbi.oussama08@gmail.com"},
		Html:    "<strong>hello world</strong>",
		Subject: "Hello from Golang",
	}

	sent, err := client.Emails.Send(params)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Failed to send email")
		return err
	}
	logger.Log.Info().Str("message_id", sent.Id).Msg("Email sent successfully")

	return nil
}
