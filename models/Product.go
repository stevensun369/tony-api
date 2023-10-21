package models

import "backend/db"

type Product struct {
  ID string `json:"ID" bson:"ID"`
  StoreID string `json:"storeID" bson:"storeID"`
  Title string `json:"title" bson:"title"`
  Desc string `json:"desc" bson:"desc"`
  Sizes []ProductSize `json:"sizes" bson:"sizes"`
  Options []ProductOption `json:"options" bson:"options"`
  Price float32 `json:"price" bson:"price"`
  Stock bool `json:"stock" bson:"stock"`
}

type ProductOption struct {
  Option string `json:"option" bson:"option"`
  Values []string `json:"values" bson:"values"`
  Stock bool `json:"stock" bson:"stock"`
}

type ProductSize struct {
  Size string `json:"size" bson:"size"`
  Price float32 `json:"price" bson:"price"`
  Stock bool `json:"stock" bson:"stock"`
}

func (p *Product) Create() error {
  p.ID = GenID(6)
  
  _, err := db.Products.InsertOne(db.Ctx, p)
  if err != nil {
    return err
  }

  return nil
}
