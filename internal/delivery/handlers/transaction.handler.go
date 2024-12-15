package handlers

import (
	"net/http"

	"github.com/damshxy/xyz-finance-app/internal/models"
	"github.com/damshxy/xyz-finance-app/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	usecase usecase.TransactionUsecase
}

func NewTransactionHandler(usecase usecase.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{
		usecase: usecase,
	}
}

func (h *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {
	var req models.Transaction
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid input",
		})
	}

	if err := h.usecase.CreateTransaction(req); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Transaction created successfully",
	})
}