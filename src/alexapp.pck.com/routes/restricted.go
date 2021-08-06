package routes

import (
	"alexapp.pck.com/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

var tokenAuth *jwtauth.JWTAuth

//InitIndexRoutes function
func InitRestrictedRoutes(c *chi.Mux, h *handler.Handler) {
	// Protected routes
	c.Group(func(r chi.Router) {
		tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))

		// Handle valid / invalid tokens. In this example, we use
		// the provided authenticator middleware, but you can write your
		// own very easily, look at the Authenticator method in jwtauth.go
		// and tweak it, its not scary.
		r.Use(jwtauth.Authenticator)

		r.Get("/restricted", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())
			name := claims["name"]
			render.Status(r, http.StatusOK)
			render.JSON(w, r, bson.M{"message": "Welcome ! " + name.(string)})
			return
		})
	})
}
