package main

import (
	"log"

	"github.com/hsrvms/fixparts/internal/server"
	"github.com/hsrvms/fixparts/pkg/config"
	"github.com/hsrvms/fixparts/pkg/db"
)

func main() {
	cfg := config.New()

	database, err := db.New(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	srv := server.New(cfg, database)
	srv.Start()
}
