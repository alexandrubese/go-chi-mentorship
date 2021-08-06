package routes

import (
	"alexapp.pck.com/handler"
	"github.com/go-chi/chi"
)

//Init Todos
func InitTodosRoutes(c *chi.Mux, h *handler.Handler) {
	c.Post("/todos/{userId}", h.CreateTodo)
	c.Get("/todos/{userId}", h.ListTodos)
	c.Get("/todo/{todoId}", h.GetTodo)
	c.Delete("/todo/{todoId}", h.DeleteTodo)
	c.Get("/todo/{todoId}/toggle", h.ToggleTodo)
}
