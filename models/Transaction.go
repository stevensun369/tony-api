package models

import "backend/db"

type Transaction struct {
  ID string `json:"ID" bson:"ID"`

  From string `json:"from" bson:"from"`
  To string `json:"to" bson:"to"`

  OrderID string `json:"orderID" bson:"orderID"`

  Hash string `json:"hash" bson:"hash"`

  Value float32 `json:"Value" bson:"Value"`
}

func (t *Transaction) Create() error {
  if (t.ID == "") {
    t.ID = GenID(12)
  }

  // wallet from out
  fromW := Wallet{}
  err := fromW.Out(t.Value)
  if err != nil {
    return err
  }

  // wallet to in
  toW := Wallet{}
  err = toW.In(t.Value)
  if err != nil {
    return err
  }

  _, err = db.Transactions.InsertOne(db.Ctx, t)

  return err
}