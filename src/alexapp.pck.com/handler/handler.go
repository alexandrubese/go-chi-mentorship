package handler

import (
	"gopkg.in/mgo.v2"
)

type (
	//Handler body
	Handler struct {
		DB *mgo.Session
	}
)

const (
	// Key (Should come from somewhere else).
	Key = "secret"
)
