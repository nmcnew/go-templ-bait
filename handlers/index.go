package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nmcnew/tournament-winner/views"
)

type IndexHandler struct {
}

func (i *IndexHandler) Register(r chi.Router) {
	r.Get("/", i.HandleIndex)
}

func (i *IndexHandler) HandleIndex(w http.ResponseWriter, r *http.Request) {
	vm := views.IndexViewModel{}
	views.Index(vm).Render(r.Context(), w)
}
