package model

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	//Feed
	Feed struct {
		ID          bson.ObjectId `json:"id" bson:"_id"`
		UserId      bson.ObjectId `json:"user_id" bson:"user_id" validate:"required"`
		Type        string        `json:"type" bson:"type" validate:"required"`
		Title       string        `json:"title" bson:"title" validate:"required"`
		Description string        `json:"description" bson:"description" validate:"required"`
		Date        string        `json:"date,omitempty" bson:"date,omitempty"`
		Location    string        `json:"location,omitempty" bson:"location,omitempty"`
		Duration    int           `json:"duration,omitempty" bson:"duration,omitempty" validate:"numeric"`
		Launched    bool          `json:"launched,omitempty" bson:"launched,omitempty"`
		Category    string        `json:"category,omitempty" bson:"category,omitempty"`
		Region      string        `json:"region,omitempty" bson:"region,omitempty"`
		Archived    bool          `json:"archived,omitempty" bson:"archived,omitempty" `
	}
)
