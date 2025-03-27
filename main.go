package main

import (
	"context"

	"github.com/ZaphCode/F-SR-ChatApp/app"
	"github.com/ZaphCode/F-SR-ChatApp/app/handlers"
)

func main() {
	server := app.New(":8080")

	// Initialize MongoDB repositories
	server.RegisterHandlers(
		handlers.NewUserHandler(),
	)

	server.Run(context.Background())
}
