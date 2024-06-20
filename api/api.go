package api

import (
	"net/http"

	"github.com/gorilla/mux"
	handler "github.com/joaocansi/essay-api/api/handler/user"
	"github.com/joaocansi/essay-api/internal/repository"
	"gorm.io/gorm"
)

type Server struct {
	db *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	return &Server{
		db: db,
	}
}

func (s *Server) Listen() {
	r := mux.NewRouter()
	subrouter := r.PathPrefix("/api/v1/").Subrouter()

	userRepository := repository.NewUserRepository(s.db)
	userHandler := handler.NewUserHandler(userRepository)
	userHandler.RegisterRoutes(subrouter)

	http.ListenAndServe(":8000", r)
}
