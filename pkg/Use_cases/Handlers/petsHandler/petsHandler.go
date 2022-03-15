package petsHandler

import "tinder_pets/pkg/Domain/response"

type PetsHandler struct {
	Repository petsRepository.Repository
}

type Handler interface {
	RegisUser(regisUser user.RegisterUser) response.Status
}
