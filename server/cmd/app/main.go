package main

import (
	"github.com/sw1pr0g/apc/server/config"
	"github.com/sw1pr0g/apc/server/internal/app"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)
}
