package models

import "fmt"

type Order struct {
  ID string `json:"ID" bson:"ID"`
  StoreID string `json:"storeID" bson:"storeID"`
  ClerkID string `json:"clerkID" bson:"clerkID"`

  Receipt []ProductConfig `json:"receipt" bson:"receipt"`

  PaymetMethod string `json:"type" bson:""`

  Value float32 `json:"value" bson:"value"`
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

func (o *Order) Build(pc []ProductConfig) {
  o.ID = GenID(12)
  o.SetValue()
  o.Receipt = pc 
}

func (o *Order) Create(ot string, storeID string, clerkID string) error {
  o.StoreID = storeID
  o.ClerkID = clerkID

  switch(ot) {
  case "loyalty": 
    fmt.Println(ot)
  default:
    fmt.Println("default")
  }

  return nil
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

