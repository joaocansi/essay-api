package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	handler "github.com/joaocansi/essay-api/cmd/api/handler/user"
	"github.com/joaocansi/essay-api/config"
	"github.com/joaocansi/essay-api/internal/repository"
	"gorm.io/gorm"
)

type APIServer struct {
	db *gorm.DB
}

func NewApiServer(db *gorm.DB) *APIServer {
	return &APIServer{
		db: db,
	}
}

func (as *APIServer) Listen() error {
	r := mux.NewRouter()
	w := r.PathPrefix("/api/v1").Subrouter()

	userRepository := repository.NewUserRepository(as.db)
	userHandler := handler.NewUserHandler(userRepository)
	userHandler.Register(w)

	log.Println("Listening on", config.Env.ServerPort)
	return http.ListenAndServe(fmt.Sprintf(":%v", config.Env.ServerPort), r)
}
