package main

import (
	"log"

	"github.com/joaocansi/essay-api/cmd/api"
	"github.com/joaocansi/essay-api/internal/database"
)

func main() {
	db, err := database.NewDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	err = api.NewApiServer(db).Listen()
	if err != nil {
		log.Fatal(err)
	}
}
