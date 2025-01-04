package main

import (
	"context"

	"below/internal/config"
	"below/internal/server"

	"github.com/gofiber/fiber/v2/log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if app, err := server.New(ctx); err != nil {
		log.Fatal(err.Error())
	} else {
		log.Fatal(app.Listen(config.Config("APP_LISTEN")))
	}
}
