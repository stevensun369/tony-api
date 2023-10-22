package models

import (
	"backend/db"

	"go.mongodb.org/mongo-driver/bson"
)

type Product struct {
  ID string `json:"ID" bson:"ID"`
  StoreID string `json:"storeID" bson:"storeID"`
  Title string `json:"title" bson:"title"`
  Desc string `json:"desc" bson:"desc"`
  Variants map[string][]ProductVariant `json:"variants" bson:"variants"`
  Options map[string][]ProductOption `json:"options" bson:"options"`
  Price float32 `json:"price" bson:"price"`
  Stock bool `json:"stock" bson:"stock"`
}

type ProductOption struct {
  Option string `json:"option" bson:"option"`
  Stock bool `json:"stock" bson:"stock"`
}

type ProductVariant struct {
  Variant string `json:"variant" bson:"variant"`
  Stock bool `json:"stock" bson:"stock"`
  Price float32 `json:"price" bson:"price"`
}

func (p *Product) Create() error {
  p.ID = GenID(6)

  _, err := db.Products.InsertOne(db.Ctx, p)
  if err != nil {
    return err
  }

  return nil
}

func (p *Product) Get(ID string) error {
  return db.Products.FindOne(
    db.Ctx,
    bson.M {
      "ID": ID,
    },
  ).Decode(&p)
}

func (p *Product) UpdateField(ID string, fieldName string, value interface{}) error {
  _, err := db.Products.UpdateOne(
    db.Ctx,
    bson.M {
      "ID": ID,
    },
    bson.M {
      "$set": bson.M {
        fieldName: value,
      },
    },
  )

  return err
}

func (p *Product) CreateVariantKey(ID string, key string) error {
  p.Get(ID)

  p.Variants[key] = []ProductVariant{}

  err := p.UpdateField(ID, "variants", p.Variants)

  return err
}

func (p *Product) RemoveVariantKey(ID string, key string) error {
  p.Get(ID)

  variants := map[string][]ProductVariant{}

  for k := range p.Variants {
    if k != key {
      variants[k] = p.Variants[k]
    }
  }

  p.Variants = variants

  err := p.UpdateField(ID, "variants", p.Variants)

  return err
}

func (p *Product) AddVariant(ID string, key string, variant ProductVariant) error {
  // getting the product
  p.Get(ID)

  // adding to the variants of a key
  p.Variants[key] = append(p.Variants[key], variant)

  // updating on the db
  err := p.UpdateField(ID, "variants", p.Variants)

  return err
}

func (p *Product) RemoveVariant(ID string, key string, variant string) error {
  // getting the product
  p.Get(ID)

  variants := []ProductVariant{}

  // adding only the variants that are not equal to the variant
  for _, v := range p.Variants[key] {
    if v.Variant != variant {
      variants = append(variants, v)
    }
  }

  p.Variants[key] = variants

  // updating on the db
  err := p.UpdateField(ID, "variants", p.Variants)

  return err
}

func (p *Product) ChangeVariant(ID string, key string, variant ProductVariant) error {
  // getting the product
  p.Get(ID)

  // changing the variant that has the same .Variant
  for i := range p.Variants[key] {
    if p.Variants[key][i].Variant == variant.Variant {
      p.Variants[key][i] = variant
    }
  }

  // updating on the db
  err := p.UpdateField(ID, "variants", p.Variants)

  return err
}

func (p *Product) CreateOptionKey(ID string, key string) error {
  p.Get(ID)

  p.Options[key] = []ProductOption{}

  err := p.UpdateField(ID, "options", p.Options)

  return err
}

func (p *Product) RemoveOptionKey(ID string, key string) error {
  p.Get(ID)

  options := map[string][]ProductOption{}

  for k := range p.Options {
    if k != key {
      options[k] = p.Options[k]
    }
  }

  p.Options = options

  err := p.UpdateField(ID, "options", p.Options)

  return err
}


func (p *Product) AddOption(ID string, key string, option ProductOption) error {
  // getting the product
  p.Get(ID)

  // adding to the options of a key
  p.Options[key] = append(p.Options[key], option)

  // updating on the db
  err := p.UpdateField(ID, "options", p.Options)

  // return nil
  return err
}

func (p *Product) RemoveOption(ID string, key string, option string) error {
  // getting the product
  p.Get(ID)

  options := []ProductOption{}

  // adding only the options that are not equal to the option
  for _, v := range p.Options[key] {
    if v.Option != option {
      options = append(options, v)
    }
  }

  p.Options[key] = options

  // updating on the db
  err := p.UpdateField(ID, "options", p.Options)

  return err
}

func (p *Product) ChangeOption(ID string, key string, option ProductOption) error {
  // getting the product
  p.Get(ID)

  // changing the option that has the same .Option
  for i := range p.Options[key] {
    if p.Options[key][i].Option == option.Option {
      p.Options[key][i] = option
    }
  }

  // updating on the db
  err := p.UpdateField(ID, "options", p.Options)

  return err
}