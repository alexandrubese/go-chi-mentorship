package routes

import (
	"alexapp.pck.com/handler"
	"github.com/go-chi/chi"
)

//InitIndexRoutes function
func InitIndexRoutes(c *chi.Mux, h *handler.Handler) {
	c.Get("/alex", h.IndexPath)
}
