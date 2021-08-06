package handler

import (
	"alexapp.pck.com/encryption"
	"alexapp.pck.com/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/render"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

//ListUsers
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	//username := c.FormValue("username")
	//password := c.FormValue("password")
	db := h.DB.Clone()
	userCtx := &model.User{}
	user := &model.User{}

	if err := render.Bind(r, userCtx); err != nil {
		render.JSON(w, r, err)
		return
	}

	if userCtx.Email != "" {
		if err := db.DB(DBName).C(UsersTable).
			Find(bson.M{"email": userCtx.Email}).
			One(&user); err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, bson.M{"Message": "Error while fetching user for patch appeared: " + err.Error()})
			return
		}
	}
	defer db.Close()

	if !checkPassword(userCtx.Password, user.Password) {
		// Throws unauthorized error
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Unathorized !"})
		return
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Email
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, bson.M{"Message": "Error while fetching user for patch appeared: " + err.Error()})
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]interface{}{
		"token": t,
		"id":    user.ID,
	})
}

func checkPassword(inputPassword string, actualPassword string) bool {
	hadedImputedPassword := encryption.CreateHash(inputPassword)
	if hadedImputedPassword == actualPassword {
		return true
	} else {
		return false
	}
}
