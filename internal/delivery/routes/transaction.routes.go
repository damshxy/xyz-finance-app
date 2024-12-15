package routes

import (
	"github.com/damshxy/xyz-finance-app/config"
	"github.com/damshxy/xyz-finance-app/database"
	"github.com/damshxy/xyz-finance-app/internal/delivery/handlers"
	"github.com/damshxy/xyz-finance-app/internal/repository"
	"github.com/damshxy/xyz-finance-app/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

func transactionRoutes(app fiber.Router) {
	cfg := config.LoadConfig()

	database, err := database.PostgresInit(cfg)
	if err != nil {
		panic("Failed to initialize database: " + err.Error())
	}

	repo := repository.NewTransactionRepository(database)
	consumerRepo := repository.NewConsumerRepository(database)
	usecase := usecase.NewTransactionUsecase(repo, consumerRepo)
	handler := handlers.NewTransactionHandler(usecase)

	app.Post("/transactions", handler.CreateTransaction)
	app.Post("/transactions/:transaction_id/refund", handler.MarkTransactionAsRefund)
	app.Get("/transactions/consumer/:consumer_id", handler.GetConsumerByID)
}