package routes

import (
	"encoding/json"
	"net/http"
	"tinder_pets/pkg/Domain/pet"
	"tinder_pets/pkg/Domain/response"
	"tinder_pets/pkg/Use_cases/Handlers/petsHandler"
	"tinder_pets/pkg/Use_cases/Helpers/responseHelper"

	"github.com/go-chi/chi"
)

type PetsRouter struct {
	Handler petsHandler.Handler
}

func (pr *PetsRouter) CreatePet(w http.ResponseWriter, r *http.Request) {
	var p pet.Pet
	token := chi.URLParam(r, "token")
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		status := response.BadRequest
		response, err := responseHelper.ResponseBuilder(status.Index(), status.String(), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500: Internal server error"))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(response))
		return
	}

	resP, status := pr.Handler.CreatePet(&p, token)
	resp, err := responseHelper.ResponseBuilder(status.Index(), status.String(), resP)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500: Internal server error"))
		return
	}
	switch status {
	case response.SuccesfulCreation:
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(resp))
		return
	case response.AlreadyExists:
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(resp))
		return
	case response.CreationFailure:
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(resp))
		return
	case response.InternalServerError, response.DBQueryError, response.DBExecutionError, response.DBLastRowIdError:
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(resp))
		return
	default:
		status = response.Unknown
		response, err := responseHelper.ResponseBuilder(status.Index(), status.String(), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500: Internal server error"))
			return
		}
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(response))
		return
	}
}

func (pr *PetsRouter) GetPetsByOwnerToken(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")
	pets, status := pr.Handler.GetPetsByOwnerToken(token)
	resp, err := responseHelper.ResponseBuilder(status.Index(), status.String(), pets)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500: Internal server error"))
		return
	}
	switch status {
	case response.SuccesfulSearch:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(resp))
		return
	case response.InternalServerError, response.DBQueryError, response.DBExecutionError, response.DBRowsError, response.DBScanError:
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(resp))
		return
	default:
		status = response.Unknown
		response, err := responseHelper.ResponseBuilder(status.Index(), status.String(), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500: Internal server error"))
			return
		}
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(response))
		return
	}
}

func (pr PetsRouter) GetPetByToken(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")
	p, status := pr.Handler.GetPetByToken(token)
	resp, err := responseHelper.ResponseBuilder(status.Index(), status.String(), p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500: Internal server error"))
		return
	}
	switch status {
	case response.SuccesfulSearch:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(resp))
		return
	case response.InternalServerError, response.DBQueryError, response.DBExecutionError, response.DBRowsError, response.DBScanError:
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(resp))
		return
	case response.NotFound:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(resp))
		return
	default:
		status = response.Unknown
		response, err := responseHelper.ResponseBuilder(status.Index(), status.String(), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500: Internal server error"))
			return
		}
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(response))
		return
	}
}

func (pr *PetsRouter) UpdatePet(w http.ResponseWriter, r *http.Request) {
	var p pet.Pet
	token := chi.URLParam(r, "token")
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		status := response.BadRequest
		response, err := responseHelper.ResponseBuilder(status.Index(), status.String(), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500: Internal server error"))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(response))
		return
	}

	status := pr.Handler.UpdatePet(token, p)
	resp, err := responseHelper.ResponseBuilder(status.Index(), status.String(), p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500: Internal server error"))
		return
	}
	switch status {
	case response.SuccesfulUpdate:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(resp))
		return
	case response.InternalServerError, response.DBQueryError, response.DBExecutionError, response.DBRowsError, response.DBScanError:
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(resp))
		return
	default:
		status = response.Unknown
		response, err := responseHelper.ResponseBuilder(status.Index(), status.String(), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500: Internal server error"))
			return
		}
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(response))
		return
	}
}

func (pr *PetsRouter) DeletePet(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")
	status := pr.Handler.DeletePet(token)
	resp, err := responseHelper.ResponseBuilder(status.Index(), status.String(), nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500: Internal server error"))
		return
	}
	switch status {
	case response.SuccesfulDelete:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(resp))
		return
	case response.InternalServerError, response.DBQueryError, response.DBExecutionError, response.DBRowsError, response.DBScanError:
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(resp))
		return
	case response.NotFound:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(resp))
		return
	default:
		status = response.Unknown
		response, err := responseHelper.ResponseBuilder(status.Index(), status.String(), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500: Internal server error"))
			return
		}
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(response))
		return
	}
}

func (pr *PetsRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Post("/owner/{token}", pr.CreatePet)

	r.Get("/owner/{token}", pr.GetPetsByOwnerToken)
	r.Get("/{token}", pr.GetPetByToken)

	r.Put("/{token}", pr.UpdatePet)

	r.Delete("/{token}", pr.DeletePet)
	return r
}
