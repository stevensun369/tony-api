package models

import "backend/db"

type Wallet struct {
  ID string `json:"ID" bson:"ID"`
  Balance float32 `json:"balance" bson:"balance"`
}

func (w Wallet) CreateWallet() error {
  w.ID = GenID(10)
  w.Balance = 0.0

  _, err := db.Wallets.InsertOne(db.Ctx, w)
  
  return err
}