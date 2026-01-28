package repository

import (
	"database/sql"
	"fmt"

	"github.com/Wahbi8/PM_Golang/DTO"
)

func GetFailedEmailsFromDB() ([]dto.EmailInfo, error) {
	var emailList []dto.EmailInfo

	db, err := sql.Open("postgres", Connection())
	if err != nil {
		return nil, fmt.Errorf("DB Connection err: %w", err)
	}
	defer db.Close()

	query := `SELECT invoice_id, type, recipient, created_at, payload, error 
			  FROM notification_logs 
			  ORDER BY created_at ASC LIMIT 5`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item dto.EmailInfo

		err := rows.Scan(
			&item.InvoiceId,
			&item.InvoiceType,
			&item.Recipient,
			&item.Created_at,
			&item.Message,
			&item.Err,
		)

		if err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}

		emailList = append(emailList, item)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return emailList, nil
}