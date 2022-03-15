package petsRepository

import (
	"os/user"
	"tinder_pets/pkg/Domain/response"
)

type Repository interface {
	RegisUser(u *user.User) response.Status
}

type UserRepository struct {
}
