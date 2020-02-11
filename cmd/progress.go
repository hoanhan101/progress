package main

import (
	"log"
	"os"
	"time"

	"github.com/hoanhan101/progress/internal/config"
	"github.com/hoanhan101/progress/internal/database"
	"github.com/hoanhan101/progress/internal/server"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	// Load the configuration.
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	// Wait for a while for the database to be ready then open its connection.
	time.Sleep(3 * time.Second)
	db, err := database.Open(cfg)
	if err != nil {
		return err
	}
	defer db.Close()

	// Bring the schema up to date.
	if err := database.Migrate(db); err != nil {
		return err
	}

	// Run the server.
	if err := server.Start(cfg, db); err != nil {
		return err
	}

	return nil
}
