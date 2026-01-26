package repository

import {
	"database/sql"
	"log"

	"github.com/Wahbi8/PM_Golang/DTO"
}

func GetFailedEmailsFromDB() {
	var emailInfo dto.EmailInfo
	db, err := sql.Open("postgres", Connection())
	if err != nil {
		log.Fatal("DB Connection err:", err)
	}
	defer db.Close()

	query := "select * from notification_logs order by created_at asec limit 5"

	_, err = db.Exec(
		query,
		emailInfo.invoice_id,
		//TODO: to be continued
	)
}