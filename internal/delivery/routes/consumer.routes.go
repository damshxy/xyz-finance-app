package routes

import (
	"github.com/damshxy/xyz-finance-app/config"
	"github.com/damshxy/xyz-finance-app/database"
	"github.com/damshxy/xyz-finance-app/internal/delivery/handlers"
	"github.com/damshxy/xyz-finance-app/internal/repository"
	"github.com/damshxy/xyz-finance-app/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

func consumerRoutes(app fiber.Router) {
	cfg := config.LoadConfig()

	database, err := database.PostgresInit(cfg)
	if err != nil {
		panic("Failed to initialize database: " + err.Error())
	}

	repo := repository.NewConsumerRepository(database)
	usecase := usecase.NewConsumerUsecase(repo)
	handler := handlers.NewConsumerHandler(usecase)

	app.Post("/consumers", handler.CreateConsumer)
	app.Get("/consumers/:id", handler.GetConsumerByID)
	app.Patch("/consumers/:id/credit", handler.UpdateConsumer)
}