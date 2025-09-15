package main

import (
	"context"

	"github.com/ZaphCode/F-SR-ChatApp/app"
	"github.com/ZaphCode/F-SR-ChatApp/app/handlers"
	"github.com/ZaphCode/F-SR-ChatApp/lib/mongodb"
	"github.com/ZaphCode/F-SR-ChatApp/repositories"
	"github.com/ZaphCode/F-SR-ChatApp/services"
	"github.com/ZaphCode/F-SR-ChatApp/utils"
)

func main() {
	db := mongodb.MustGetMongoClient(utils.MONGO_DEV_URI).Database("fsr-sandbox")

	// * Repositories
	userRepository := repositories.NewMongoDBUserRepository(db.Collection("users"))

	//* Services
	userService := services.NewUserService(userRepository)

	server := app.New(8080)

	// * Handlers
	server.RegisterHandlers(
		handlers.NewAuthHandler(userService),
	)

	server.OnShutdown(func() {
		db.Client().Disconnect(context.TODO())
	})

	server.Run(context.Background())
}
