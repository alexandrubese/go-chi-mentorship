package model

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	//Recipe
	Recipe struct {
		ID           bson.ObjectId `json:"id" bson:"_id"`
		UserId       bson.ObjectId `json:"user_id" bson:"user_id" validate:"required"`
		Title        string        `json:"title" bson:"title" validate:"required"`
		Ingredients  []string      `json:"ingredients" bson:"ingredients" validate:"required"`
		Instructions string        `json:"instructions" bson:"instructions" validate:"required"`
	}
)
