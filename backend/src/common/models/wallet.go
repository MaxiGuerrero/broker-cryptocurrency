package models

type Wallet struct {
	UserId string `json:"user_id" bson:"user_id"`
	Funds  []Fund `json:"funds" bson:"funds"`
}
