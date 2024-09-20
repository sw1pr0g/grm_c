package main

import (
	"github.com/sw1pr0g/amazing-people-calendar/server/config"
	"github.com/sw1pr0g/amazing-people-calendar/server/internal/app"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
