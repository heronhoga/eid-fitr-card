package utils

import "github.com/oklog/ulid/v2"

func GenerateCardID() string {
	cardId := ulid.Make().String()
	return cardId
}