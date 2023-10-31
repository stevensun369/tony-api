package models

import (
	"backend/db"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

type Wallet struct {
  ID string `json:"ID" bson:"ID"`
  Balance int `json:"balance" bson:"balance"`
}

func (w *Wallet) CreateWallet() error {
  w.ID = GenID(10)
  w.Balance = 0.0

  _, err := db.Wallets.InsertOne(db.Ctx, w)
  
  return err
}

func (w *Wallet) Get(ID string) error {
  err := db.Wallets.FindOne(db.Ctx, bson.M{
    "ID": ID,
  }).Decode(&w)

  return err
}

func (w *Wallet) Out(amount int) error {
  w.Get(w.ID)

  if (w.Balance > amount) {
    _, err := db.Wallets.UpdateOne(
      db.Ctx,
      bson.M {"ID": w.ID},
      bson.M {
        "$set": bson.M {
          "balance": w.Balance - amount,
        },
      },
    )

    if err != nil {
      return err
    }

    w.Balance = w.Balance - amount
  
    return nil
  } else {
    return errors.New("fonduri insuficiente")
  }
}

func (w *Wallet) In(amount int) error {
  w.Get(w.ID)

  _, err := db.Wallets.UpdateOne(
    db.Ctx,
    bson.M {"ID": w.ID},
    bson.M {
      "$set": bson.M {
        "balance": w.Balance + amount,
      },
    },
  )

  if err != nil {
    return err
  }

  w.Balance = w.Balance + amount

  return nil
}