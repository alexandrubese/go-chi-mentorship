package initscripts

import (
	"alexapp.pck.com/handler"
	"alexapp.pck.com/routes"
	"github.com/go-chi/chi"
)

//InitAppRoutes body
func InitAppRoutes(c *chi.Mux, h *handler.Handler) {
	routes.InitIndexRoutes(c, h)
	routes.InitRecipesRoutes(c, h)
	routes.InitRestrictedRoutes(c, h)
	routes.InitUserRoutes(c, h)
	routes.InitFeedRoutes(c, h)
	routes.InitTodosRoutes(c, h)
}
