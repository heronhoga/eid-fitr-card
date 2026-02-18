package models

type CardRequest struct {
	CardID string `json:"card_id"`
}

type CardResponse struct {
	CardID string `json:"card_id"`
	To string `json:"to"`
	From string `json:"from"`
	Title string `json:"title"`
	Description string `json:"description"`
}