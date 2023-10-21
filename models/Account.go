package models

import (
	"backend/db"
	"backend/env"
	"errors"

	sj "github.com/brianvoe/sjwt"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
  ID string `json:"ID" bson:"ID"`
  UserName string `json:"username" bson:"username"`
  Phone string `json:"phone" bson:"phone"`
  WalletID string `json:"wallet" bson:"wallet"`
  Password string `json:"password" bson:"password"`
}

func (u *User) GenUserToken() (string, error) {
  claims, err := sj.ToClaims(u)
  claims.SetExpiresAt(expirationTime)

  token := claims.Generate(env.JWTKey)
  return token, err
}

func (u *User) ParseUserToken(token string) error {
  verified := sj.Verify(token, []byte(env.JWTKey))

  if !verified {
    return nil
  }

  claims, err := sj.Parse(token)
  if err != nil {
    return err
  }

  err = claims.Validate()
  claims.ToStruct(&u)

  return err
} 

func (u *User) CreateUser(phone string) error {
  id := GenID(8)

  var wallet Wallet
  wallet.CreateWallet()

  u.ID = id
  u.Phone = phone
  u.WalletID = wallet.ID

  _, err := db.Users.InsertOne(db.Ctx, u)
  return err
}

func (u *User) AddUsername(ID string, username string) (error) {
  tempUser := User{}
  err := db.Users.FindOne(db.Ctx, bson.M{
    "username": username,
  }).Decode(&tempUser)
  
  if err == nil {
    return errors.New("numele de utilizator există deja")
  }

  err = db.Users.FindOneAndUpdate(
    db.Ctx,
    bson.M{
      "ID": ID,
    }, 
    bson.M{
      "$set": bson.M{
        "username": username,
      },
    },
  ).Decode(&u)
  u.UserName = username

  if err != nil {
    return errors.New("nu s-a putut găsi")
  }

  return nil
}

func (u *User) AddPassword(ID string, password string) (error) {
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)

  if err != nil {
    return err
  }

  err = db.Users.FindOneAndUpdate(
    db.Ctx,
    bson.M{
      "ID": ID,
    },
    bson.M{
      "$set": bson.M{
        "password": hashedPassword,
      },
    },
  ).Decode(&u)

  return err
}