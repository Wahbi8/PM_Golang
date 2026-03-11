package repository

import (
	"database/sql"
	"time"

	"github.com/Wahbi8/PM_Golang/DTO"
	"github.com/Wahbi8/PM_Golang/logger"
)

func InsertFailedMsgs(emailInfo *dto.EmailInfo, errorMsg string) {
	db, err := sql.Open("postgres", Connection())
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("DB connection failed")
	}
	defer db.Close()

	query := `insert into notification_logs (invoice_id, type, recipient, created_at, payload, error)
				values($1, $2, $3, $4, $5, $6)`

	_, err = db.Exec(
		query,
		emailInfo.InvoiceId,
		emailInfo.InvoiceType,
		emailInfo.Recipient,
		time.Now(),
		emailInfo.Message,
		errorMsg,
	)

	if err != nil {
		logger.Log.Error().Err(err).Msg("Failed to insert failed message")
	}
}
