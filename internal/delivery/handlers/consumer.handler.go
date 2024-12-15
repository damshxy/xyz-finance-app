package handlers

import (
	"net/http"

	"github.com/damshxy/xyz-finance-app/internal/models"
	"github.com/damshxy/xyz-finance-app/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type ConsumerHandler struct {
	usecase usecase.ConsumerUsecase
}

func NewConsumerHandler(usecase usecase.ConsumerUsecase) *ConsumerHandler {
	return &ConsumerHandler{
		usecase: usecase,
	}
}

func (h *ConsumerHandler) CreateConsumer(c *fiber.Ctx) error {
	var req models.Consumer
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid input",
		})
	}

	if err := h.usecase.CreateConsumer(req); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Consumer created successfully",
	})
}

func (h *ConsumerHandler) GetConsumerByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	consumer, err := h.usecase.GetConsumerByID(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Consumer not found",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Consumer retrieved successfully",
		"consumer": consumer,
	})
}

func (h *ConsumerHandler) UpdateConsumer(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	var req struct {
		NewCreditLimit float64 `json:"new_credit_limit"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	if err := h.usecase.UpdateConsumer(id, req.NewCreditLimit); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Consumer updated successfully",
	})
}