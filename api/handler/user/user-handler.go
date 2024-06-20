package handler

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/joaocansi/essay-api/internal/repository"
	"github.com/joaocansi/essay-api/utils"
)

type Handler struct {
	userRepository repository.UserRepository
}

func NewUserHandler(userRepository repository.UserRepository) *Handler {
	return &Handler{userRepository: userRepository}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/users", h.CreateUser).Methods("POST")
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user CreateUserDTO
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("body is missing"))
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid fields %v", errors))
		return
	}

	if _, err := h.userRepository.GetUserByEmail(user.Email); err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", user.Email))
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("internal server error"))
		return
	}

	user.Password = string(hashedPassword)
	err = h.userRepository.CreateUser(user.ToModel())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
