package routes

import (
	"alexapp.pck.com/handler"
	"github.com/go-chi/chi"
)

//Init Recipes
func InitRecipesRoutes(c *chi.Mux, h *handler.Handler) {
	c.Post("/recipes/{userId}", h.CreateRecipe)
	c.Get("/recipes/{userId}", h.ListRecipes)
	c.Get("/recipe/{recipeId}", h.GetRecipe)
	c.Delete("/recipes/{id}", h.DeleteRecipe)
}
