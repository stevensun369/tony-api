package models

import "backend/db"

type Transaction struct {
  ID string `json:"ID" bson:"ID"`
  From string `json:"from" bson:"from"`
  To string `json:"to" bson:"to"`
  Hash string `json:"hash" bson:"hash"`

  Receipt []ProductConfig `json:"receipt" bson:"receipt"`

  Value float32 `json:"value" bson:"value"`
}

type ProductConfig struct {
  ProductID string `json:"productID" bson:"productID"`

  Title string `json:"title" bson:"title"`
  Variants []ProductVariant `json:"variant" bson:"variant"`
  Options []ProductOption `json:"options" bson:"options"`

  Price float32 `json:"price" bson:"price"`
  Quantity int `json:"quantity" bson:"quantity"`
}

func (t *Transaction) SetValue() {
  for _, pc := range t.Receipt {
    t.Value += pc.GetPrice()
  }
} 

func (t *Transaction) Create() error {
  t.ID = GenID(12)

  _, err := db.Transactions.InsertOne(db.Ctx, t)
  return err
}

func (p *ProductConfig) GetPrice() float32 {
  var variantsValue float32 = 0 

  for _, v := range p.Variants {
    variantsValue += v.Price
  }

  return  float32(p.Quantity) * (p.Price + variantsValue)
}

