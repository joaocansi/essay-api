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

type AuthHandler struct {
	userRepository *repository.UserRepository
}

func NewAuthHandler(userRepository *repository.UserRepository) *AuthHandler {
	return &AuthHandler{
		userRepository: userRepository,
	}
}

func (ah *AuthHandler) Register(r *mux.Router) {
	r.HandleFunc("/users/auth", ah.AuthenticateUser).Methods("POST")
}

func (ah *AuthHandler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var data AuthenticateUserDTO
	if err := utils.ToJSON(r.Body, &data); err != nil {
		httpres.New(w).InternalServerError()
		return
	}

	u, err := ah.userRepository.GetUserByEmail(data.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			httpres.New(w).Error(http.StatusUnauthorized, "email/password is incorrect")
			return
		}
		httpres.New(w).InternalServerError()
		return
	}

	if err := utils.ComparePassword(u.Password, data.Password); err != nil {
		httpres.New(w).Error(http.StatusUnauthorized, "email/password is incorrect")
		return
	}

	token, err := utils.GenerateJwt(u.ID)
	if err != nil {
		httpres.New(w).InternalServerError()
		return
	}

	httpres.New(w).Send(http.StatusCreated, map[string]string{"token": token})
}
