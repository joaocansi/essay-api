package handler

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joaocansi/essay-api/cmd/api/httpres"
	"github.com/joaocansi/essay-api/internal/repository"
	"github.com/joaocansi/essay-api/utils"
	"gorm.io/gorm"
)

type UserHandler struct {
	userRepository *repository.UserRepository
}

func NewUserHandler(userRepository *repository.UserRepository) *UserHandler {
	return &UserHandler{
		userRepository: userRepository,
	}
}

func (uh *UserHandler) Register(r *mux.Router) {
	r.HandleFunc("/users", uh.CreateUser)
}

func (uh *UserHandler) Test(w http.ResponseWriter, r *http.Request) {

}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user CreateUserDTO

	if err := utils.ToJSON(r.Body, &user); err != nil {
		httpres.New(w).InternalServerError()
		return
	}

	_, err := uh.userRepository.GetUserByEmail(user.Email)
	if err == nil {
		httpres.New(w).Error(http.StatusConflict, "email already registered")
		return
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		httpres.New(w).InternalServerError()
		return
	}

	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		httpres.New(w).InternalServerError()
		return
	}

	if err = uh.userRepository.CreateUser(user.ToModel()); err != nil {
		httpres.New(w).InternalServerError()
		return
	}

	httpres.New(w).Send(http.StatusCreated, map[string]string{})
}
