package repositories

import (
	"context"

	"github.com/heronhoga/eid-fitr-card/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type CardRepository struct {
	collection *mongo.Collection
}

func NewCardRepository(db *mongo.Database) *CardRepository {
	return &CardRepository{
		collection: db.Collection("cards"),
	}
}

func (r *CardRepository) CreateCard(ctx context.Context, request *models.Card) (string, error) {

	_, err := r.collection.InsertOne(ctx, request)
	if err != nil {
		return "", err
	}

	return request.CardID, nil
}

func (r *CardRepository) GetCard(ctx context.Context, request models.GetCardRequest) (models.GetCardResponse, error) {

	var response models.GetCardResponse

	filter := bson.M{"cardid": request.CardID}

	err := r.collection.FindOne(ctx, filter).Decode(&response)
	if err != nil {
		return models.GetCardResponse{}, err
	}

	return response, nil
}

