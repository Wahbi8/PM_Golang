package repository

func Connection() string {
	//TODO: read connection string from .env
	connStr := "postgres://postgres:Postgresqlaccount1@localhost:5432/PM_logs?sslmode=disable"
	return connStr
}
	