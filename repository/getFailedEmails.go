package repository

import (
	"database/sql"
	"log"

	"github.com/Wahbi8/PM_Golang/DTO"
)

func GetFailedEmailsFromDB() dto.EmailInfo {
	//TODO: get list not just one ------- to correct the code
	var emailInfo dto.EmailInfo
	db, err := sql.Open("postgres", Connection())
	if err != nil {
		log.Fatal("DB Connection err:", err)
	}
	defer db.Close()

	query := `select invoice_id, type, recipient, created_at, payload, error 
				from notification_logs order by created_at asc limit 5`

	_, err = db.Query(
		query,
		emailInfo.InvoiceId,
		emailInfo.InvoiceType,
		emailInfo.Recipient,
		emailInfo.Created_at,
		emailInfo.Message,
		emailInfo.Err,
	)

	if err != nil {
		log.Fatal(err)
	}

	return emailInfo
}