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

//Create Todo
func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	myValidator := &CustomValidator{validator: validator.New()}

	isValidObjectId := bson.IsObjectIdHex(userId)
	if !isValidObjectId {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Please supply a valid objectid string"})
		return
	}
	// Bind
	todo := &model.Todo{ID: bson.NewObjectId(), UserId: bson.ObjectIdHex(userId)}

	//for added funct render.Bind can be used (and create a new Bind function for that Object to do after the bind is done
	if err := json.NewDecoder(r.Body).Decode(todo); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message2": "Error appeared2: " + err.Error()})
		return
	}

	if err := myValidator.Validate(todo); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message1": "Error appeared1: " + err.Error()})
		return
	}

	// Save todo
	db := h.DB.Clone()
	defer db.Close()

	if err := db.DB(DBName).C(TodosTable).Insert(todo); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message3": "Error appeared3: " + err.Error()})
		return
	}
	defer db.Close()

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, todo)
}

//List Todos
func (h *Handler) ListTodos(w http.ResponseWriter, r *http.Request) {
	// Retrieve user based on id from database
	userId := chi.URLParam(r, "userId")

	isValidObjectId := bson.IsObjectIdHex(userId)
	if !isValidObjectId {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Please supply a valid objectid string"})
		return
	}

	todos := []*model.Todo{}
	db := h.DB.Clone()
	if err := db.DB(DBName).C(TodosTable).
		Find(bson.M{"user_id": bson.ObjectIdHex(userId)}).
		All(&todos); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Error appeared: " + err.Error()})
		return
	}
	defer db.Close()

	render.Status(r, http.StatusOK)
	render.JSON(w, r, todos)
}

//Get Todos
func (h *Handler) GetTodo(w http.ResponseWriter, r *http.Request) {
	todoId := chi.URLParam(r, "todoId")

	isValidObjectId := bson.IsObjectIdHex(todoId)
	if !isValidObjectId {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Please supply a valid objectid string"})
		return
	}

	todo := &model.Todo{ID: bson.ObjectIdHex(todoId)}
	db := h.DB.Clone()
	if err := db.DB(DBName).C(TodosTable).
		FindId(todo.ID).One(&todo); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Error appeared: " + err.Error()})
		return
	}

	defer db.Close()

	render.Status(r, http.StatusOK)
	render.JSON(w, r, todo)
}

//Delete Todo
func (h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	// Retrieve user based on id from database
	id := chi.URLParam(r, "todoId")

	isValidObjectId := bson.IsObjectIdHex(id)
	if !isValidObjectId {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Please supply a valid objectid string"})
		return
	}

	todo := &model.Todo{ID: bson.ObjectIdHex(id)}
	db := h.DB.Clone()
	if err := db.DB(DBName).C(TodosTable).
		RemoveId(todo.ID); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Error appeared: " + err.Error()})
		return
	}

	defer db.Close()

	render.Status(r, http.StatusOK)
	render.JSON(w, r, bson.M{"Message": "Todo with id: " + id + " was deleted"})
}

//Toggle Todo
func (h *Handler) ToggleTodo(w http.ResponseWriter, r *http.Request) {
	// Retrieve user based on id from database
	id := chi.URLParam(r, "todoId")

	isValidObjectId := bson.IsObjectIdHex(id)
	if !isValidObjectId {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Please supply a valid objectid string"})
		return
	}

	todo := &model.Todo{ID: bson.ObjectIdHex(id)}
	db := h.DB.Clone()
	if err := db.DB(DBName).C(TodosTable).
		FindId(todo.ID).
		One(&todo); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Error while fetching todo for toggle appeared: " + err.Error()})
		return
	}

	todo.Checked = !todo.Checked

	if err := db.DB(DBName).C(TodosTable).
		UpdateId(todo.ID, todo); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Error appeared: " + err.Error()})
		return
	}

	defer db.Close()

	render.Status(r, http.StatusOK)
	render.JSON(w, r, todo)
}
