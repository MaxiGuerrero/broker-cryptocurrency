package database

import (
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ObjectIDToString(value interface{}) string {
	objectId, ok := value.(primitive.ObjectID)
	if !ok {
		log.Panicln("Cannot cast objectId. Invalid type")
	}
	return objectId.Hex()
}

func DateTimeToTime(value interface{}) time.Time {
	if value == nil {
		return time.Time{}
	}
	t, ok := value.(primitive.DateTime)
	if !ok {
		return time.Time{}
	}
	return t.Time()
}
