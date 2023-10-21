package db

import (
	"log"

	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

var Accounts *mongo.Collection

func GetCollection(col string, client *mongo.Client) (*mongo.Collection) {
  return client.Database("dev").Collection(col)
}

func InitDB(MongoURI string) {
  var err error

  Client, err = mongo.Connect(
    context.Background(),
    options.Client().ApplyURI(MongoURI),
  )

  if err != nil {
    log.Fatal(err)
  }

  Accounts = GetCollection("accounts", Client)
}