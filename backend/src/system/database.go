package system

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	client *mongo.Client
}

var DbConnection = os.Getenv("DB_URI_CONNECTION")
var DbName = os.Getenv("DB_NAME")

func NewDatabase(ctxParent context.Context) *Database {
	ctx, cancel := context.WithTimeout(ctxParent, 10*time.Second)
	defer cancel()
	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(DbConnection))
	if err != nil {
		log.Fatalln(err.Error())
	}
	// Send a ping to confirm a successful connection
	if err := cli.Database(DbName).RunCommand(ctxParent, bson.D{primitive.E{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}
	log.Printf("Pinged your database %v. You successfully connected to MongoDB!", DbName)
	return &Database{
		client: cli,
	}
}

func (d *Database) GetConnector(collection string) *mongo.Collection {
	if d.client != nil {
		log.Fatalln("Database client is not initialized, call NewDatabase() in main.go")
	}
	return d.client.Database(DbName).Collection(collection)
}
