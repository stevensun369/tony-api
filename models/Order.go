package models

import (
	"backend/db"
	"backend/utils"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type Order struct {
  ID string `json:"ID" bson:"ID"`
  StoreID string `json:"storeID" bson:"storeID"`
  ClerkID string `json:"clerkID" bson:"clerkID"`

  Receipt []ProductConfig `json:"receipt" bson:"receipt"`

  TransactionID string `json:"transactionID" bson:"transactionID"`

  Value int `json:"value" bson:"value"`

  // date time
  Date string `json:"date" bson:"date"`
  Time string `json:"time" bson:"time"`
}

type ProductConfig struct {
  ProductID string `json:"productID" bson:"productID"`

  Title string `json:"title" bson:"title"`
  Variants map[string]ShortProductVariant `json:"variants" bson:"variants"`
  Options map[string]string `json:"options" bson:"options"`

  Price int `json:"price" bson:"price"`
  Quantity int `json:"quantity" bson:"quantity"`
}

type ShortProductVariant struct {
  Variant string `json:"variant" bson:"variant"`
  Price int `json:"price" bson:"price"`
}

func GetOrders(filter interface{}, sort interface{}) ([]Order, error) {
  orders := []Order {}

  opts := options.Find().SetSort(sort)

  cursor, err := db.Orders.Find(db.Ctx, filter, opts)
  if err != nil {
    return orders, err
  }

  err = cursor.All(db.Ctx, &orders)
  if err != nil {
    return orders, err
  }

  return orders, nil
}

func (o *Order) Create(ot string, pc []ProductConfig, storeID string, clerkID string, ID string) error {
  o.Receipt = pc 
  o.StoreID = storeID
  o.ClerkID = clerkID
  o.ID = GenID(12)
  o.Date = utils.GetToday()
  o.Time = utils.GetNow()

  o.SetValue()

  switch(ot) {
  case "loyalty": 
    cashback := Transaction{
      From: storeID,
      To: ID,
      Value: utils.ApplyCashbackRate(o.Value),
      OrderID: o.ID,
      Tag: "cashback",
      CreatedAt: time.Now(),
    }
    
    if err := cashback.Create(); err != nil {
      return err
    }
  case "app":
    fmt.Println("app")
  default:
    fmt.Println("default")
  }

  _, err := db.Orders.InsertOne(db.Ctx, o)

  return err
}

func (o *Order) SetValue() {
  for _, pc := range o.Receipt {
    o.Value += pc.GetPrice()
  }
}

func (p *ProductConfig) GetPrice() int {
  var variantsValue int = 0 

  for _, v := range p.Variants {
    variantsValue += v.Price
  }

  return  int(p.Quantity) * (p.Price + variantsValue)
}

