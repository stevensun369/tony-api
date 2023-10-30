package models

import (
	"backend/db"
	"backend/utils"
	"fmt"
	"time"
)

type Order struct {
  ID string `json:"ID" bson:"ID"`
  StoreID string `json:"storeID" bson:"storeID"`
  ClerkID string `json:"clerkID" bson:"clerkID"`

  Receipt []ProductConfig `json:"receipt" bson:"receipt"`

  TransactionID string `json:"transactionID" bson:"transactionID"`

  Value float32 `json:"value" bson:"value"`

  // date time
  Date string `json:"date" bson:"date"`
  Time string `json:"time" bson:"time"`
}

type ProductConfig struct {
  ProductID string `json:"productID" bson:"productID"`

  Title string `json:"title" bson:"title"`
  Variants []ShortProductVariant `json:"variants" bson:"variants"`
  Options []string `json:"options" bson:"options"`

  Price float32 `json:"price" bson:"price"`
  Quantity int `json:"quantity" bson:"quantity"`
}

type ShortProductVariant struct {
  Variant string `json:"variant" bson:"variant"`
  Price float32 `json:"price" bson:"price"`
}

type PaymentMethod struct {
  Type string `json:"type" bson:"type"`
  Reference string `json:""`
}

func GetOrders(filter interface{}) ([]Order, error) {
  orders := []Order {}

  cursor, err := db.Orders.Find(db.Ctx, filter)
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

func (p *ProductConfig) GetPrice() float32 {
  var variantsValue float32 = 0 

  for _, v := range p.Variants {
    variantsValue += v.Price
  }

  return  float32(p.Quantity) * (p.Price + variantsValue)
}

