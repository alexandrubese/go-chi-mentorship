package handler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"

	"alexapp.pck.com/encryption"
	"alexapp.pck.com/model"
	"gopkg.in/mgo.v2/bson"
)

//ListUsers  thisis
func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	/*userID := userIDFromToken(c)
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	// Defaults
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 100
	}*/

	// Retrieve users from database
	users := []*model.User{}
	db := h.DB.Clone()
	if err := db.DB(DBName).C(UsersTable).
		Find(bson.M{}).
		All(&users); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Error appeared: " + err.Error()})
		return
	}
	defer db.Close()

	render.Status(r, http.StatusOK)
	render.JSON(w, r, users)
}

//GetUser  thisis
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	// Retrieve user based on id from database
	id := chi.URLParam(r, "id")

	isValidObjectId := bson.IsObjectIdHex(id)
	if !isValidObjectId {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Please supply a valid objectid string"})
		return
	}

	user := &model.User{ID: bson.ObjectIdHex(id)}
	db := h.DB.Clone()
	if err := db.DB(DBName).C(UsersTable).
		FindId(user.ID).
		One(&user); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Error appeared: " + err.Error()})
		return
	}

	defer db.Close()

	render.Status(r, http.StatusOK)
	render.JSON(w, r, user)
}

//CreateUser  thisis
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Bind
	u := &model.User{ID: bson.NewObjectId()}

	//for added funct render.Bind can be used (and create a new Bind function for that Object to do after the bind is done
	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		return
	}

	// Validate
	if u.Email == "" || u.Password == "" {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "invalid email or password"})
		return
	}

	///Hash password
	hashPassword(u)

	// Save user
	db := h.DB.Clone()
	if err := db.DB(DBName).C(UsersTable).Insert(u); err != nil {
		return
	}
	defer db.Close()
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, u)
}

//UpdateUser body
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Retrieve user based on id from database
	id := chi.URLParam(r, "id")

	isValidObjectId := bson.IsObjectIdHex(id)
	if !isValidObjectId {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Please supply a valid objectid string"})
		return
	}

	user := &model.User{ID: bson.ObjectIdHex(id)}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Error appeared: " + err.Error()})
		return
	}

	//Hash password
	hashPassword(user)

	db := h.DB.Clone()
	if err := db.DB(DBName).C(UsersTable).
		UpdateId(user.ID, user); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Error appeared: " + err.Error()})
		return
	}

	defer db.Close()

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, user)
}

//UpdatePartUser body
func (h *Handler) UpdatePartUser(w http.ResponseWriter, r *http.Request) {
	// Retrieve user based on id from database
	id := chi.URLParam(r, "id")
	db := h.DB.Clone()

	isValidObjectId := bson.IsObjectIdHex(id)
	if !isValidObjectId {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Please supply a valid objectid string"})
		return
	}

	user := &model.User{ID: bson.ObjectIdHex(id)}
	body := &model.User{ID: bson.ObjectIdHex(id)}

	if err := db.DB(DBName).C(UsersTable).
		FindId(user.ID).
		One(&user); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Error while fetching user for patch appeared: " + err.Error()})
		return
	}

	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Cannot bind user data to patch data: " + err.Error()})
		return
	}

	if body.Password != "" {
		//Hash password need an if here to not hash the password all the time
		hashPassword(body)
	}
	//Assign the body data to the user object
	user = body

	if err := db.DB(DBName).C(UsersTable).
		UpdateId(bson.ObjectIdHex(id), user); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Error while patching user: " + err.Error()})
		return
	}

	defer db.Close()

	render.Status(r, http.StatusOK)
	render.JSON(w, r, user)
}

//DeleteUser body
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Retrieve user based on id from database
	id := chi.URLParam(r, "id")

	isValidObjectId := bson.IsObjectIdHex(id)
	if !isValidObjectId {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Please supply a valid objectid string"})
		return
	}

	user := &model.User{ID: bson.ObjectIdHex(id)}
	db := h.DB.Clone()
	if err := db.DB(DBName).C(UsersTable).
		RemoveId(user.ID); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Error appeared: " + err.Error()})
		return
	}

	defer db.Close()

	render.Status(r, http.StatusOK)
	render.JSON(w, r, user)
}

//hashPassword hashes the password of a User
func hashPassword(user *model.User) *model.User {
	if user.Password != "" {
		user.Password = encryption.CreateHash(user.Password)
	}
	return user
}
