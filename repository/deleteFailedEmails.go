package repository

import (
	"database/sql"
	"log"
	
	"github.com/google/uuid"
)

func DeleteFailedEmail(invoiceID uuid.UUID) {
    db, err := sql.Open("postgres", Connection())
	if err != nil {
		log.Fatal("DB Connection err:", err)
	}
    defer db.Close()
    db.Exec("DELETE FROM notification_logs WHERE invoice_id = $1", invoiceID)
}