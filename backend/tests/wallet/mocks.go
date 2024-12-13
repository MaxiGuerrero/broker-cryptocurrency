package wallet

import "backend/src/common/models"

const userId = "1"

var userWallet = models.Wallet{
	UserId: userId,
	Funds: []models.Fund{{
		Currency: "ARS",
		Amount:   2,
	}},
}
