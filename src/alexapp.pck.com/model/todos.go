package model

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	//Recipe
	Todo struct {
		ID      bson.ObjectId `json:"id" bson:"_id"`
		UserId  bson.ObjectId `json:"user_id" bson:"user_id" validate:"required"`
		Title   string        `json:"title" bson:"title" validate:"required"`
		Details string        `json:"details" bson:"details" validate:"required"`
		Checked bool          `json:"checked" bson:"checked" default:"false"`
		DueDate string        `json:"due_date" bson:"due_date" validate:"required"`
	}
)
