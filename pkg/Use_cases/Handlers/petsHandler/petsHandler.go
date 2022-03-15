package petsHandler

import (
	"tinder_pets/internal/data/Infrastructure/petsRepository"
	"tinder_pets/pkg/Domain/pet"
	"tinder_pets/pkg/Domain/response"

	"github.com/gofrs/uuid"
)

type PetsHandler struct {
	Repository petsRepository.Repository
}

type Handler interface {
	CreatePet(p *pet.Pet, ownerToken string) response.Status
	GetPetsByOwnerToken(ownertoken string) ([]pet.Pet, response.Status)
	GetPetByToken(token string) (interface{}, response.Status)
	UpdatePet(token string, p pet.Pet) response.Status
	DeletePet(token string) response.Status
}

func (ph *PetsHandler) CreatePet(p *pet.Pet, ownerToken string) response.Status {
	token, _ := uuid.NewV4()
	p.Token = token.String()

	p.Ownertoken = ownerToken
	return ph.Repository.CreatePet(p, ownerToken)
}

func (ph *PetsHandler) GetPetsByOwnerToken(ownerToken string) ([]pet.Pet, response.Status) {
	return ph.Repository.GetPetsByOwnerToken(ownerToken)
}

func (ph *PetsHandler) GetPetByToken(token string) (interface{}, response.Status) {
	return ph.Repository.GetPetByToken(token)
}

func (ph *PetsHandler) UpdatePet(token string, p pet.Pet) response.Status {
	return ph.Repository.UpdatePet(token, p)
}

func (ph *PetsHandler) DeletePet(token string) response.Status {
	return ph.Repository.DeletePet(token)
}
