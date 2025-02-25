package models

import (
	"backend/db"
	"backend/env"
	"errors"

	sj "github.com/brianvoe/sjwt"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	ID       string `json:"ID" bson:"ID"`
	WalletID string `json:"walletID" bson:"walletID"`
	UserName string `json:"username" bson:"username"`
	Phone    string `json:"phone" bson:"phone"`
}

func (u *User) GenToken() (string, error) {
	claims, err := sj.ToClaims(&u)
	claims.SetExpiresAt(expirationTime)

	token := claims.Generate(env.JWTKey)
	return token, err
}

func (u *User) ParseToken(token string) error {
	verified := sj.Verify(token, []byte(env.JWTKey))

	if !verified {
		return nil
	}

	claims, err := sj.Parse(token)
	if err != nil {
		return err
	}

	err = claims.Validate()
	if err != nil {
		return err
	}
	err = claims.ToStruct(&u)

	return err
}

func (u *User) Create(phone string) error {
	id := GenID(8)

	var wallet Wallet
	wallet.CreateWallet()

	u.ID = id
	u.Phone = phone
	u.WalletID = wallet.ID

	_, err := db.Users.InsertOne(db.Ctx, u)
	return err
}

func (u *User) Get(filter interface{}) error {
	err := db.Users.FindOne(db.Ctx, filter).Decode(&u)

	return err
}

func UserCheck(filter interface{}) bool {
	u := User{}
	err := u.Get(filter)

	if err == nil {
		return true
	} else {
		return false
	}
}

func (u *User) AddUsername(ID string, username string) error {
	exists := UserCheck(bson.M{
		"username": username,
	})

	if exists {
		return errors.New("numele de utilizator există deja")
	}

	err := db.Users.FindOneAndUpdate(
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

	if err != nil {
		return errors.New("nu s-a putut găsi")
	}

	u.UserName = username

	return nil
}
