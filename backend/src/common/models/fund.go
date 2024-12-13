package models

type Fund struct {
	Currency Currency `json:"currency" bson:"currency"`
	Amount   float64  `json:"amount" bson:"amount"`
}
