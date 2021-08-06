package handler

import (
	"github.com/go-chi/render"
	"net/http"
)

//IndexPath
func (h *Handler) IndexPath(w http.ResponseWriter, r *http.Request) {

	responseJSON := &JSResp{Msg: "Hello from Alex!"}
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, responseJSON)
}
