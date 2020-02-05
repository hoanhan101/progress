package main

import (
	"log"

	"github.com/hoanhan101/progress/internal/config"
	"github.com/hoanhan101/progress/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	s := server.NewServer(cfg)
	s.Run()
}
