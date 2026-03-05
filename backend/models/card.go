package models

type Card struct {
	CardID string `json:"card_id"`
	To string `json:"to"`
	From string `json:"from"`
	Title string `json:"title"`
	Description string `json:"description"`
}
type GetCardRequest struct {
	CardID string `json:"card_id"`
}

type GetCardResponse struct {
	CardID string `json:"card_id"`
	To string `json:"to"`
	From string `json:"from"`
	Title string `json:"title"`
	Description string `json:"description"`
}

type CreateCardRequest struct {
	To string `json:"to"`
	From string `json:"from"`
	Title string `json:"title"`
	Description string `json:"description"`
}