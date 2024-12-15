package usecase

import (
	"errors"

	"github.com/damshxy/xyz-finance-app/internal/models"
	"github.com/damshxy/xyz-finance-app/internal/repository"
)

type ConsumerUsecase interface {
	CreateConsumer(consumer models.Consumer) error
	GetConsumerByID(id int) (models.Consumer, error)
	UpdateConsumer(id int, newCreditLimit float64) error
}

type consumerUsecase struct {
	repo repository.ConsumerRepository
}

func NewConsumerUsecase(repo repository.ConsumerRepository) ConsumerUsecase {
	return &consumerUsecase{
		repo: repo,
	}
}

func (u *consumerUsecase) CreateConsumer(consumer models.Consumer) error {
	// Validate required fields
	if consumer.NIK == "" || consumer.FullName == "" {
		return errors.New("NIK and FullName are required")
	}

	// Check if NIK already exists
	exists, err := u.repo.ConsumerExistByNIK(consumer.NIK)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("NIK already exists")
	}

	// Set default credit limit for a new consumer
	if consumer.CreditLimit == 0 {
		consumer.CreditLimit = 1000000.0
	}

	return u.repo.CreateConsumer(&consumer)
}

func (u *consumerUsecase) GetConsumerByID(id int) (models.Consumer, error) {
	consumer, err := u.repo.GetConsumerByID(id)
	if err != nil {
		return consumer, errors.New("consumer not found")
	}

	return consumer, nil
}

func (u *consumerUsecase) UpdateConsumer(id int, newCreditLimit float64) error {
	consumer, err := u.repo.GetConsumerByID(id)
	if err != nil {
		return errors.New("consumer not found")
	}

	if newCreditLimit < 0 {
		return errors.New("credit limit can't be negative")
	}

	consumer.CreditLimit = newCreditLimit
	return u.repo.UpdateConsumer(&consumer)
}