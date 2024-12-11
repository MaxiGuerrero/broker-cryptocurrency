package authentication

import (
	"backend/src/authentication/models"
	"backend/src/system/database"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository struct {
	database *database.Database
}

const usersCollection = "users"

func (a *AuthRepository) CreateUser(username, password, email string) *models.UserInfo {
	result, err := a.database.GetCollection(usersCollection).InsertOne(context.TODO(), models.User{
		Username:  username,
		Password:  password,
		Email:     email,
		Status:    models.ACTIVE,
		Role:      models.USER,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		log.Panicf("Error on create user document: %v", err.Error())
	}
	log.Printf("User %v has been registered", username)
	return &models.UserInfo{
		UserId:   database.ObjectIDToString(result.InsertedID),
		Username: username,
		Email:    email,
		Role:     models.USER.String(),
		Status:   models.ACTIVE.String(),
	}
}

func (a *AuthRepository) FindUserByUsername(username string) *models.User {
	var result bson.M
	filter := bson.M{"username": username}
	err := a.database.GetCollection(usersCollection).FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		log.Panicf("Error on get user by username or email: %v", err.Error())
	}
	return &models.User{
		ID:        database.ObjectIDToString(result["_id"]),
		Username:  result["username"].(string),
		Password:  result["password"].(string),
		Email:     result["email"].(string),
		Status:    models.Status(result["status"].(int32)),
		Role:      models.Role(result["role"].(int32)),
		CreatedAt: database.DateTimeToTime(result["created_at"]),
		UpdatedAt: database.DateTimeToTime(result["updated_at"]),
		DeletedAt: database.DateTimeToTime(result["deleted_at"]),
	}
}
