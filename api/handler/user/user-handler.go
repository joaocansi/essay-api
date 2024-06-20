package handler

import (
	"net/http"

	"github.com/joaocansi/essay-api/internal/repository"
	"github.com/joaocansi/essay-api/pkg"
)

type Handler struct {
	UserRepository repository.UserRepository
	Mux            *http.ServeMux
}

func NewHandler(userRepository repository.UserRepository) *Handler {
	return &Handler{UserRepository: userRepository}
}

func (h *Handler) RegisterRoutes() {
	h.Mux.HandleFunc("/user", h.CreateUser)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	pkg.WriteError(w, http.StatusInternalServerError, nil)
}
