package repository

import (
	"database/sql"
	"github.com/google/uuid"
)

func DeleteFailedEmail(invoiceID uuid.UUID) {
    db, _ := sql.Open("postgres", Connection())
    defer db.Close()
    db.Exec("DELETE FROM notification_logs WHERE invoice_id = $1", invoiceID)
}