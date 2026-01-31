package dto

import (
	"time"

	"github.com/google/uuid"
)

type EmailInfo struct{
    Sender string			
    Recipient string			`json:"RecipientEmail"`
    Message string				`json:"Body"`
    Subject string				`json:"subject"`
    InvoiceId uuid.UUID			`json:"InvoiceId"`
    UserId uuid.UUID
    MessageType MessageType
    InvoiceType InvoiceType		`json:"InvoiceType"`
	Retry int					`json:"retry"`
    Created_at time.Time
    Err string
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