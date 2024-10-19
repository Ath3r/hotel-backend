package server

import (
	"fmt"

	"github.com/Ath3r/hotel-backend/internal/config"
	"github.com/Ath3r/hotel-backend/internal/db"
	"github.com/Ath3r/hotel-backend/internal/handlers"
	"github.com/Ath3r/hotel-backend/internal/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type App struct {
	Server *fiber.App
}

var serverConfig = fiber.Config{
	ErrorHandler: middlewares.ErrorHandler,
}

func NewApp() (*App, error) {

	// Setup database
	client, err := db.ConnectMongo()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Setup router
	app, err := setupRouter()

	if err != nil {
		return nil, err
	}

	api := app.Group("/api")
	v1 := api.Group("/v1")

	userHandler := handlers.NewUserHandler(db.NewMongoUserStore(client))

	v1.Get("/users/:id", userHandler.HandleGetUser)
	v1.Get("/users", userHandler.HandleGetUsers)
	v1.Post("/users", userHandler.HandlePostUser)

	return &App{
		Server: app,
	}, nil
}



func (a *App) Start() error {
	fmt.Printf("Starting server on port %d\n", config.AppConfig.Port)
	return a.Server.Listen(fmt.Sprintf(":%d", config.AppConfig.Port))
}

func setupRouter() (*fiber.App, error) {
	app := fiber.New(serverConfig)

	app.Use(cors.New())
	app.Use(logger.New())

	return app, nil
}
