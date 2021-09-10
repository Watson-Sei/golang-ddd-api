package handler

import (
	"encoding/json"
	"fmt"
	"golang-ddd-api/app/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler interface {
	Index(http.ResponseWriter, *http.Request)
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserhandler(uu usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: uu,
	}
}

func (uh userHandler) Index(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["name"]

	user, err := uh.userUseCase.Search(name)
	if err != nil {
		http.Error(w, fmt.Sprintf("...: %w", err), 400)
	}

	u, err := json.Marshal(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("...: %w", err), 400)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(u)
}
