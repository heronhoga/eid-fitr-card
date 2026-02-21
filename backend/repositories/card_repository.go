package repositories

import (
	"github.com/gofiber/fiber/v3"
	"github.com/heronhoga/eid-fitr-card/models"
	"github.com/jackc/pgx/v5"
)

type CardRepository struct {
	db *pgx.Conn
}

func NewCardRepository(db *pgx.Conn) *CardRepository {
	return &CardRepository{db: db}
}

func (r *CardRepository) CreateCard(context fiber.Ctx, request *models.Card) (string, error) {
	var cardID string

	err := r.db.QueryRow(
		context,
		`INSERT INTO cards ("card_id", "to", "from", title, description, type)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING card_id`,
		request.CardID,
		request.To,
		request.From,
		request.Title,
		request.Description,
		request.Type,
	).Scan(&cardID)

	if err != nil {
		return "", err
	}

	return cardID, nil
}

func (r *CardRepository) GetCard(context fiber.Ctx, request models.GetCardRequest) (models.GetCardResponse, error) {
	var response models.GetCardResponse

	err := r.db.QueryRow(
		context,
		`SELECT "card_id", "to", "from", title, description, type
		FROM cards
		WHERE card_id = $1`,
		request.CardID,
	).Scan(
		&response.CardID,
		&response.To,
		&response.From,
		&response.Title,
		&response.Description,
		&response.Type,
	)

	if err != nil {
		return models.GetCardResponse{}, err
	}

	return response, nil
}

