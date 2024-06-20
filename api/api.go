package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joaocansi/essay-api/internal/config"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
}

func NewServer(DB *gorm.DB) *Server {
	return &Server{
		DB: DB,
	}
}

func (s *Server) Listen() {
	r := mux.NewRouter()
	r.PathPrefix("/api/v1/")

	http.ListenAndServe(fmt.Sprintf(":%v", config.Env.Api.Port), r)
}
