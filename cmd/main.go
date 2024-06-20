package main

import (
	"github.com/joaocansi/essay-api/api"
	"github.com/joaocansi/essay-api/internal/config"
	"github.com/joaocansi/essay-api/internal/database"
)

func main() {
	DB := database.NewDatabase(config.Env.DB)
	server := api.NewServer(DB)
	server.Listen()
}
