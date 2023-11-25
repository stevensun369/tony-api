package models

import (
	"backend/db"
	"backend/env"
	"errors"

	sj "github.com/brianvoe/sjwt"
	"go.mongodb.org/mongo-driver/bson"
)

type StoreAdmin struct {
	ID      string `json:"ID" bson:"ID"`
	StoreID string `json:"storeID" bson:"storeID"`
}

func (sa *StoreAdmin) GenToken() (string, error) {
	claims, err := sj.ToClaims(&sa)
	claims.SetExpiresAt(expirationTime)

	token := claims.Generate(env.JWTKey)
	return token, err
}

func (sa *StoreAdmin) ParseToken(token string) error {
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
	err = claims.ToStruct(&sa)

	return err
}

func (sa *StoreAdmin) Create(ID string) error {
	if !UserCheck(bson.M{"ID": ID}) {
		return errors.New("nu există utilizator")
	}

	sa.ID = ID

	_, err := db.StoreAdmins.InsertOne(db.Ctx, &sa)

	return err
}

func (sa *StoreAdmin) Get(ID string) error {
	err := db.StoreAdmins.FindOne(db.Ctx,
		bson.M{
			"ID": ID,
		},
	).Decode(&sa)

	return err
}

func (sa *StoreAdmin) Check(ID string) bool {
	err := sa.Get(ID)

	if err == nil {
		return true
	} else {
		return false
	}
}

func (sa *StoreAdmin) AddStore(ID string, storeID string) error {
	s := Store{}

	if !s.Check(storeID) {
		return errors.New("nu există vânzătorul")
	}

	err := db.StoreAdmins.FindOneAndUpdate(
		db.Ctx,
		bson.M{
			"ID": ID,
		},
		bson.M{
			"storeID": storeID,
		},
	).Decode(&sa)

	sa.StoreID = storeID

	return err
}
