package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Instanciamos los handlers de los endpoints
func New() http.Handler {
	r := chi.NewRouter()

	ur := &UserRouter{
		Handler: &userHandler.UserHandler{
			Repository: &userRepository.UserRepository{},
		},
	}

	r.Mount("/pets", ur.Routes())

	//Retornamos la api ya construida
	return r

}
