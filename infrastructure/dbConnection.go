package infrastructure

func Connection() string {
	connStr := "postgres://postgres:Postgresqlaccount1@localhost:5432/PM_logs?sslmode=disable"
	return connStr
}
