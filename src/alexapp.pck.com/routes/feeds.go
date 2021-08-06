package routes

import (
	"alexapp.pck.com/handler"
	"github.com/go-chi/chi"
)

//InitUserRoutes function
func InitFeedRoutes(c *chi.Mux, h *handler.Handler) {
	c.Post("/feeds/{userId}", h.CreateFeed)
	c.Get("/feeds/{userId}", h.ListFeeds)
	c.Delete("/feeds/{id}", h.DeleteFeed)
}
