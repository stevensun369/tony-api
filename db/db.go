package db

import (
	"fmt"
	"log"

	"context"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var RedisClient *redis.Client
var Ctx = context.Background()

var Users *mongo.Collection
var Wallets *mongo.Collection

var Stores *mongo.Collection
var StoreAdmins *mongo.Collection
var Clerks *mongo.Collection

var Products *mongo.Collection
var Transactions *mongo.Collection

func GetCollection(col string, client *mongo.Client) (*mongo.Collection) {
  return client.Database("dev").Collection(col)
}

func InitDB(MongoURI string) {
  var err error

  MongoClient, err = mongo.Connect(
    context.Background(),
    options.Client().ApplyURI(MongoURI),
  )

  if err != nil {
    log.Fatal(err)
  }

  Users = GetCollection("users", MongoClient)
  Wallets = GetCollection("wallets", MongoClient)

  Stores = GetCollection("stores", MongoClient)
  StoreAdmins = GetCollection("storeAdmins", MongoClient)
  Clerks = GetCollection("clerks", MongoClient)
  
  Products = GetCollection("products", MongoClient)
  Transactions = GetCollection("transactions", MongoClient)

  fmt.Println("connected to MongoDB")
}

func InitCache(RedisOptions *redis.Options) {
  RedisClient = redis.NewClient(RedisOptions)

  pong, _ := RedisClient.Ping(Ctx).Result()
  if pong == "PONG" {
    fmt.Println("connected to Redis")
  } else {
    fmt.Println("not connected to redis")
  }
}

func Set(key string, value string) error {
  err := RedisClient.Set(Ctx, key, value, 0).Err()

  return err
}

func Get(key string) (string, error) {
  val, err := RedisClient.Get(Ctx, key).Result()

  return val, err
}

func Del(key string) error {
  _, err := RedisClient.Del(Ctx, key).Result()

  return err
}