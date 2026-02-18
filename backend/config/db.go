package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func NewDBConn() (*pgx.Conn, error) {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, dbName)
	conn, err := pgx.Connect(context.Background(), connString)
	if err!=nil{
		return nil,err
	}
	return conn, nil
}