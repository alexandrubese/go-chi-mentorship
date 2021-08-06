package routes

import (
	"alexapp.pck.com/handler"
	"github.com/go-chi/chi"
)

//InitUserRoutes function
func InitUserRoutes(c *chi.Mux, h *handler.Handler) {
	//Unguarded routes
	c.Post("/users/login", h.Login)
	c.Post("/users", h.CreateUser)

	//Guarded routes
	c.Get("/users", h.ListUsers)
	c.Get("/users/{id}", h.GetUser)
	c.Put("/users/{id}", h.UpdateUser)
	c.Patch("/users/{id}", h.UpdatePartUser)
	c.Delete("/users/{id}", h.DeleteUser)
}
