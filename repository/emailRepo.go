package repository

import (
	"database/sql"
	"log"

	"github.com/Wahbi8/PM_Golang/DTO"
)

func InsertFailedMsgs(emailInfo *dto.EmailInfo) {
	db, err := sql.Open("postgres", Connection())
	if err != nil {
		log.Fatal("DB Connection err:", err)
	}
	defer db.Close()

	query := 	`insert into notification_logs (invoice_id, type, recipient, created_at, payload, error)
				values($1, $2, $3, $4, $5, $6)`
	
	_, err = db.Exec(
		query,
		
	)
}