package handler

import (
	"alexapp.pck.com/model"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

//CreateUser  thisis
func (h *Handler) CreateFeed(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	myValidator := &CustomValidator{validator: validator.New()}

	isValidObjectId := bson.IsObjectIdHex(userId)
	if !isValidObjectId {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Please supply a valid objectid string"})
		return
	}
	// Bind
	feed := &model.Feed{ID: bson.NewObjectId(), UserId: bson.ObjectIdHex(userId)}

	//for added funct render.Bind can be used (and create a new Bind function for that Object to do after the bind is done
	if err := json.NewDecoder(r.Body).Decode(feed); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Error appeared: " + err.Error()})
		return
	}

	if err := myValidator.Validate(feed); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Error appeared: " + err.Error()})
		return
	}

	// Save user
	db := h.DB.Clone()
	defer db.Close()
	if err := db.DB(DBName).C(FeedsTable).Insert(feed); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Error appeared: " + err.Error()})
		return
	}
	defer db.Close()

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, feed)
}

//GetUser  thisis
func (h *Handler) ListFeeds(w http.ResponseWriter, r *http.Request) {
	// Retrieve user based on id from database
	userId := chi.URLParam(r, "userId")

	isValidObjectId := bson.IsObjectIdHex(userId)
	if !isValidObjectId {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Please supply a valid objectid string"})
		return
	}

	feeds := []*model.Feed{}
	db := h.DB.Clone()
	if err := db.DB(DBName).C(FeedsTable).
		Find(bson.M{"user_id": bson.ObjectIdHex(userId)}).
		All(&feeds); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Error appeared: " + err.Error()})
		return
	}
	defer db.Close()

	render.Status(r, http.StatusOK)
	render.JSON(w, r, feeds)
}

//DeleteUser body
func (h *Handler) DeleteFeed(w http.ResponseWriter, r *http.Request) {
	// Retrieve user based on id from database
	id := chi.URLParam(r, "id")

	isValidObjectId := bson.IsObjectIdHex(id)
	if !isValidObjectId {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Please supply a valid objectid string"})
		return
	}

	feed := &model.User{ID: bson.ObjectIdHex(id)}
	db := h.DB.Clone()
	if err := db.DB(DBName).C(FeedsTable).
		RemoveId(feed.ID); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Error appeared: " + err.Error()})
		return
	}

	defer db.Close()

	render.Status(r, http.StatusOK)
	render.JSON(w, r, bson.M{"Message": "Feed with id: " + id + " was deleted"})
}
