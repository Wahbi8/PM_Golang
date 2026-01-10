
package Services

import (
    "github.com/resend/resend-go/v2"
    "fmt"
    "os"
    "log"
    "github.com/joho/godotenv"
    "github.com/google/uuid"
)

type EmailInfo struct{
    Sender string
    Recipient string
    Message string
    Subject string
    InvoiceId uuid.UUID
    UserId uuid.UUID
    MessageType MessageType
    InvoiceType InvoiceType
}
type MessageType int

const (
    Email MessageType = iota
    SMS
)

type InvoiceType int

const (
    draft InvoiceType = iota
    sent
    paid
    canceled
)


func SendEmail() {
    godotenv.Load()

    apiKey := os.Getenv("Resend_api_key")
    if apiKey == "" {
        log.Fatal("can't load the api key")
    }

    client := resend.NewClient(apiKey)

    params := &resend.SendEmailRequest{
        From:    "Acme <onboarding@resend.dev>",
        To:      []string{"wahbi.oussama08@gmail.com"},
        Html:    "<strong>hello world</strong>",
        Subject: "Hello from Golang",
        // Cc:      []string{"cc@example.com"},
        // Bcc:     []string{"bcc@example.com"},
        // ReplyTo: "replyto@example.com",
    }

    sent, err := client.Emails.Send(params)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Println(sent.Id)
}