package routes

import (
	"net/http"
	"tinder_pets/internal/data/Infrastructure/petsRepository"
	"tinder_pets/pkg/Use_cases/Handlers/petsHandler"

	"github.com/go-chi/chi"
)

// Instanciamos los handlers de los endpoints
func New() http.Handler {
	r := chi.NewRouter()

	ur := &PetsRouter{
		Handler: &petsHandler.PetsHandler{
			Repository: &petsRepository.PetsRepository{},
		},
	}

	r.Mount("/pets", ur.Routes())

	//Retornamos la api ya construida
	return r

}
