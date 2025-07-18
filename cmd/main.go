//go:build wireinject
// +build wireinject

package main

import (
	"coding-games/config"
	"coding-games/infrastructure/server"
	"coding-games/internal/core/service/lucksvc"
	"coding-games/internal/handler/gamehdl"
	"log"

	"github.com/gofiber/fiber/v3"åå
	"github.com/google/wire"
)

type dependencies struct {
	server *fiber.App
}

func initialize() *dependencies {
	wire.Build(
		server.New,
		gamehdl.New,
		lucksvc.New,
		wire.Struct(new(dependencies), "server"),
	)
	return nil
}

func main() {
	d := initialize()

	log.Printf("Server is running on port %s", config.Get().Server.Port)

	// Graceful shutdown and other server logic can be added here
	defer func() {
		if err := d.server.Shutdown(); err != nil {
			log.Printf("Error during shutdown: %s", err.Error())
		}
		log.Println("Server gracefully stopped")
	}()

	// Start the server on the configured port
	if err := d.server.Listen(":" + config.Get().Server.Port); err != nil {
		log.Fatalf("Failed to start server: %s", err.Error())
	}
}
