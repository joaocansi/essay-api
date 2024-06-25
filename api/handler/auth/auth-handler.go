package handler

import (
	"fmt"
	"net/http"

	"github.com/joaocansi/essay-api/internal/repository"
	"github.com/joaocansi/essay-api/utils"
)

type AuthHandler struct {
	userRepository *repository.UserRepository
}

func NewAuthHandler(userRepository *repository.UserRepository) *AuthHandler {
	return &AuthHandler{
		userRepository: userRepository,
	}
}

type Token struct {
	token string
}

func (ah *AuthHandler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var login AuthenticateUserDTO
	if err := utils.ParseJSON(r, &login); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user, err := ah.userRepository.GetUserByEmail(login.Email)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("email/password is incorrect"))
		return
	}

	if err = utils.ComparePassword(user.Password, login.Password); err != nil {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("email/password is incorrect"))
		return
	}

	token, err := utils.GenerateJwt(user.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	utils.WriteJSON(w, http.StatusAccepted, Token{
		token: token,
	})
}
