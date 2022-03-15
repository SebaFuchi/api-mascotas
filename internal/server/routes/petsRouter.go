package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

type PetsRouter struct {
	Handler petsHandler.Handler
}

func (pr *PetsRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Post("/", pr.CreatePet)

	return r
}
