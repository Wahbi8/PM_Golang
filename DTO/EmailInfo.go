package dto

import (
	"time"

	"github.com/google/uuid"
)

type EmailInfo struct{
    Sender string			
    Recipient string			`json:"to"`
    Message string				`json:"body"`
    Subject string				`json:"subject"`
    InvoiceId uuid.UUID			`json:"invoice_id"`
    UserId uuid.UUID
    MessageType MessageType
    InvoiceType InvoiceType		`json:"invoice_type"`
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