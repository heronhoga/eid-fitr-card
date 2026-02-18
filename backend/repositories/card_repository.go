package repositories

import "github.com/jackc/pgx/v5"

type CardRepository struct {
	db *pgx.Conn
}

func NewCardRepository(db *pgx.Conn) *CardRepository {
	return &CardRepository{db: db}
}

