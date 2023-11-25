package models

import (
	"backend/db"

	"go.mongodb.org/mongo-driver/bson"
)

type Store struct {
	ID   string   `json:"ID" bson:"ID"`
	Name string   `json:"name" bson:"name"`
	Open string   `json:"open" bson:"open"`
	Tags []string `json:"tags" bson:"tags"`
}

func (s *Store) Get(ID string) error {
	return db.Stores.FindOne(db.Ctx, bson.M{
		"ID": ID,
	}).Decode(&s)
}

func (s *Store) Check(ID string) bool {
	return s.Get(ID) == nil
}

func (s *Store) Change(ID string, update interface{}) error {
	return db.Stores.FindOneAndUpdate(db.Ctx,
		bson.M{
			"ID": ID,
		},
		bson.M{
			"$set": update,
		},
	).Decode(&s)
}

func (s *Store) ChangeName(ID string, name string) error {
	err := s.Change(ID, bson.M{
		"name": name,
	})

	s.Name = name

	return err
}

func (s *Store) ChangeOpen(ID string, open string) error {
	err := s.Change(ID, bson.M{
		"open": open,
	})
	s.Open = open

	return err
}
