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

//Create Recipe
func (h *Handler) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	myValidator := &CustomValidator{validator: validator.New()}

	isValidObjectId := bson.IsObjectIdHex(userId)
	if !isValidObjectId {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Please supply a valid objectid string"})
		return
	}
	// Bind
	recipe := &model.Recipe{ID: bson.NewObjectId(), UserId: bson.ObjectIdHex(userId)}

	//for added funct render.Bind can be used (and create a new Bind function for that Object to do after the bind is done
	if err := json.NewDecoder(r.Body).Decode(recipe); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message2": "Error appeared2: " + err.Error()})
		return
	}

	if err := myValidator.Validate(recipe); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message1": "Error appeared1: " + err.Error()})
		return
	}

	// Save user
	db := h.DB.Clone()
	defer db.Close()
	if err := db.DB(DBName).C(RecipesTable).Insert(recipe); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message3": "Error appeared3: " + err.Error()})
		return
	}
	defer db.Close()

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, recipe)
}

//List Recipes
func (h *Handler) ListRecipes(w http.ResponseWriter, r *http.Request) {
	// Retrieve user based on id from database
	userId := chi.URLParam(r, "userId")

	isValidObjectId := bson.IsObjectIdHex(userId)
	if !isValidObjectId {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Please supply a valid objectid string"})
		return
	}

	recipes := []*model.Recipe{}
	db := h.DB.Clone()
	if err := db.DB(DBName).C(RecipesTable).
		Find(bson.M{"user_id": bson.ObjectIdHex(userId)}).
		All(&recipes); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Error appeared: " + err.Error()})
		return
	}
	defer db.Close()

	render.Status(r, http.StatusOK)
	render.JSON(w, r, recipes)
}

//Get Recipes
func (h *Handler) GetRecipe(w http.ResponseWriter, r *http.Request) {
	recipeId := chi.URLParam(r, "recipeId")

	isValidObjectId := bson.IsObjectIdHex(recipeId)
	if !isValidObjectId {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Please supply a valid objectid string"})
		return
	}

	recipe := &model.Recipe{ID: bson.ObjectIdHex(recipeId)}
	db := h.DB.Clone()
	if err := db.DB(DBName).C(RecipesTable).
		FindId(recipe.ID).One(&recipe); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Error appeared: " + err.Error()})
		return
	}

	defer db.Close()

	render.Status(r, http.StatusOK)
	render.JSON(w, r, recipe)
}

//Delete Recipe
func (h *Handler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	// Retrieve user based on id from database
	id := chi.URLParam(r, "id")

	isValidObjectId := bson.IsObjectIdHex(id)
	if !isValidObjectId {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Please supply a valid objectid string"})
		return
	}

	recipe := &model.Recipe{ID: bson.ObjectIdHex(id)}
	db := h.DB.Clone()
	if err := db.DB(DBName).C(RecipesTable).
		RemoveId(recipe.ID); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Error appeared: " + err.Error()})
		return
	}

	defer db.Close()

	render.Status(r, http.StatusOK)
	render.JSON(w, r, bson.M{"Message": "Recipe with id: " + id + " was deleted"})
}
