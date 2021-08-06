package initscripts

import (
	"gopkg.in/mgo.v2"
	"log"
)

//InitDbConnection body
func InitDbConnection() *mgo.Session {
	//db, err := mgo.Dial("mongodb") //docker container hostname
	db, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal("Database connection problem: " + err.Error())
	}
	return db
}
