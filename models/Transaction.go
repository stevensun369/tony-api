package models

import (
	"backend/db"
	"errors"
	"time"
)

type Transaction struct {
  ID string `json:"ID" bson:"ID"`
  Tag string `json:"tag" bson:"tag"`

  From string `json:"from" bson:"from"`
  To string `json:"to" bson:"to"`

  OrderID string `json:"orderID" bson:"orderID"`

  Hash string `json:"hash" bson:"hash"`

  Value int `json:"value" bson:"value"`
  CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}

func (t *Transaction) Create() error {
  if (t.ID == "") {
    t.ID = GenID(12)
  }

  // wallet from out
  fromW := Wallet{ID: t.From}
  err := fromW.Out(t.Value)
  if err != nil {
    return errors.New("from")
  }

  // wallet to in
  toW := Wallet{ID: t.To}
  err = toW.In(t.Value)
  if err != nil {
    return errors.New("to")
  }

  _, err = db.Transactions.InsertOne(db.Ctx, t)

  return err
}