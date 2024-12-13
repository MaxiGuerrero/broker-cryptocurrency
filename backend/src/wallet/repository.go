package wallet

import (
	commonModels "backend/src/common/models"
	"backend/src/system/database"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type WalletRepository struct {
	Database *database.Database
}

const walletCollection = "wallets"

func (w *WalletRepository) CreateWallet(userId string) {
	funds := []commonModels.Fund{}
	currencies := commonModels.GetCurrencies()
	for _, currency := range currencies {
		funds = append(funds, commonModels.Fund{
			Currency: currency,
			Amount:   0,
		})
	}
	_, err := w.Database.GetCollection(walletCollection).InsertOne(context.TODO(), commonModels.Wallet{
		UserId: userId,
		Funds:  funds,
	})
	if err != nil {
		log.Panicf("Error on create wallet for user: %v", err.Error())
	}
}

func (w *WalletRepository) GetWalletByUserId(userId string) *commonModels.Wallet {
	var result commonModels.Wallet
	filter := bson.M{"user_id": userId}
	err := w.Database.GetCollection(walletCollection).FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		log.Panicf("Error on get wallet by username: %v", err.Error())
	}
	return &commonModels.Wallet{
		UserId: result.UserId,
		Funds:  result.Funds,
	}
}
