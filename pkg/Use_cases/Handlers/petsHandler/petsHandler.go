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
	CreatePet(p *pet.Pet, ownertoken string) response.Status
	GetPetsByOwnerToken(ownertoken string) ([]pet.Pet, response.Status)
	GetPetByToken(token string) (pet.Pet, response.Status)
	UpdatePet(token string, p pet.Pet) response.Status
	DeletePet(token string) response.Status
}

func (ph *PetsHandler) CreatePet(p *pet.Pet, ownertoken string) response.Status {
	token, _ := uuid.NewV4()
	p.Token = token.String()

	return ph.Repository.CreatePet(p, ownertoken)
}

func (ph *PetsHandler) GetPetsByOwnerToken(ownertoken string) ([]pet.Pet, response.Status) {
	return ph.Repository.GetPetsByOwnerToken(ownertoken)
}

func (ph *PetsHandler) GetPetByToken(token string) (pet.Pet, response.Status) {
	return ph.Repository.GetPetByToken(token)
}

func (ph *PetsHandler) UpdatePet(token string, p pet.Pet) response.Status {
	return ph.Repository.UpdatePet(token, p)
}

func (ph *PetsHandler) DeletePet(token string) response.Status {
	return ph.Repository.DeletePet(token)
}
