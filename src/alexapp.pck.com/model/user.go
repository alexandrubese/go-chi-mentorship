package model

import (
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type (
	//User sdfasf
	User struct {
		ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Email    string        `json:"email" bson:"email"`
		Password string        `json:"password,omitempty" bson:"password"`
		Token    string        `json:"token,omitempty" bson:"-"`
	}
)

func (u *User) Bind(r *http.Request) error {
	return nil
}
