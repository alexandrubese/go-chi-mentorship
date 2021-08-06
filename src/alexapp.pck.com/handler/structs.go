package handler

import "gopkg.in/go-playground/validator.v9"

const (
	DBName              = "goApi"
	UsersTable   string = "users"
	FeedsTable   string = "feeds"
	RecipesTable string = "recipes"
	TodosTable   string = "todos"
)

//An example of doing a struct of values
/*
var Tables = struct {
	UsersTable string
	PostsTable string
}{
	UsersTable: "users",
	PostsTable: "posts",
}
*/

//JSResp asd
type JSResp struct {
	Msg string `json:"Msg"`
}

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
