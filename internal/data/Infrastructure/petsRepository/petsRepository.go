package petsRepository

import (
	"tinder_pets/pkg/Domain/pet"
	"tinder_pets/pkg/Domain/response"
	"tinder_pets/pkg/Use_cases/Helpers/dbHelper"
)

type Repository interface {
	CreatePet(p *pet.Pet, ownerToken string) response.Status
	GetPetsByOwnerToken(ownertoken string) ([]pet.Pet, response.Status)
	GetPetByToken(token string) (interface{}, response.Status)
	UpdatePet(token string, p pet.Pet) response.Status
	DeletePet(token string) response.Status
}

type PetsRepository struct {
}

func (pr *PetsRepository) CreatePet(p *pet.Pet, ownerToken string) response.Status {
	sqlCon, err := dbHelper.GetDB()
	if err != nil {
		return response.InternalServerError
	}
	defer sqlCon.Close()

	selForm, err := sqlCon.Prepare("SELECT token FROM pets WHERE name = ?")
	if err != nil {
		return response.DBQueryError
	}
	defer selForm.Close()

	res, err := selForm.Query(p.Name)
	if err != nil {
		return response.DBQueryError
	}

	defer res.Close()

	if res.Next() {
		return response.AlreadyExists
	}

	insForm, err := sqlCon.Prepare("INSERT INTO pets(token, owner_token, name, type, sex, image) VALUES (?,?,?,?,?,?)")
	if err != nil {
		return response.DBQueryError
	}
	defer insForm.Close()

	result, err := insForm.Exec(
		p.Token,
		p.OwnerToken,
		p.Name,
		p.Type,
		p.Sex,
		p.Image,
	)
	if err != nil {
		return response.DBExecutionError

	} else {
		rows, err := result.RowsAffected()
		if err != nil {
			return response.DBRowsError
		}
		if rows == 0 {
			return response.CreationFailure
		}
	}
	return response.SuccesfulCreation
}

func (pr *PetsRepository) GetPetsByOwnerToken(ownertoken string) ([]pet.Pet, response.Status) {
	sqlCon, err := dbHelper.GetDB()
	if err != nil {
		return nil, response.InternalServerError
	}
	defer sqlCon.Close()

	selfForm, err := sqlCon.Prepare("SELECT token, owner_token, name, type, sex, image FROM pets WHERE owner_token = ?")
	if err != nil {
		return nil, response.DBQueryError
	}
	defer selfForm.Close()

	rows, err := selfForm.Query(ownertoken)
	if err != nil {
		return nil, response.DBQueryError
	}
	defer rows.Close()

	var pets []pet.Pet
	for rows.Next() {
		var p pet.Pet
		err = rows.Scan(
			&p.Token,
			&p.OwnerToken,
			&p.Name,
			&p.Type,
			&p.Sex,
			&p.Image,
		)
		if err != nil {
			return nil, response.DBScanError
		}

		pets = append(pets, p)
	}
	return pets, response.SuccesfulSearch
}

func (pr *PetsRepository) GetPetByToken(token string) (interface{}, response.Status) {
	sqlCon, err := dbHelper.GetDB()
	if err != nil {
		return nil, response.InternalServerError
	}
	defer sqlCon.Close()
	var p pet.Pet

	selForm, err := sqlCon.Prepare("SELECT token, owner_token, name, type, sex, image FROM pets WHERE token = ?")
	if err != nil {
		return nil, response.DBQueryError
	}
	defer selForm.Close()

	result, err := selForm.Query(token)
	if err != nil {
		return nil, response.DBQueryError
	}
	defer result.Close()

	if result.Next() {
		err = result.Scan(
			&p.Token,
			&p.OwnerToken,
			&p.Name,
			&p.Type,
			&p.Sex,
			&p.Image,
		)
		if err != nil {
			return nil, response.DBScanError
		}

		return p, response.SuccesfulSearch
	}
	return nil, response.NotFound
}

func (pr *PetsRepository) UpdatePet(token string, p pet.Pet) response.Status {
	sqlCon, err := dbHelper.GetDB()
	if err != nil {
		return response.InternalServerError
	}
	defer sqlCon.Close()

	updForm, err := sqlCon.Prepare("UPDATE pets SET name = ?, type = ?, sex = ?, image = ? WHERE token = ?")
	if err != nil {
		return response.DBQueryError
	}

	result, err := updForm.Exec(
		p.Name,
		p.Type,
		p.Sex,
		p.Image,
		token,
	)
	if err != nil {
		return response.DBExecutionError

	} else {
		_, err := result.RowsAffected()
		if err != nil {
			return response.DBRowsError
		}
	}
	return response.SuccesfulUpdate
}

func (pr *PetsRepository) DeletePet(token string) response.Status {
	sqlCon, err := dbHelper.GetDB()
	if err != nil {
		return response.InternalServerError
	}
	defer sqlCon.Close()

	delForm, err := sqlCon.Prepare("DELETE FROM pets WHERE token = ?")
	if err != nil {
		return response.DBQueryError
	}
	defer delForm.Close()

	result, err := delForm.Exec(token)
	if err != nil {
		return response.DBExecutionError

	} else {
		rows, err := result.RowsAffected()
		if err != nil {
			return response.DBRowsError
		}
		if rows == 0 {
			return response.NotFound
		}
	}
	return response.SuccesfulDelete
}
