package usecase

import (
	"errors"
	"fmt"

	"github.com/damshxy/xyz-finance-app/internal/models"
	"github.com/damshxy/xyz-finance-app/internal/repository"
)

type TransactionUsecase interface {
	CreateTransaction(transaction models.Transaction) error
}

type transactionUsecase struct {
	repo repository.TransactionRepository
	consumerRepo repository.ConsumerRepository
}

func NewTransactionUsecase(repo repository.TransactionRepository, consumerRepo repository.ConsumerRepository) TransactionUsecase {
	return &transactionUsecase{
		repo: repo,
		consumerRepo: consumerRepo,
	}
}

func (u *transactionUsecase) CreateTransaction(transaction models.Transaction) error {
	if transaction.ConsumerID == 0 || transaction.OTR <= 0 {
		return errors.New("invalid transaction data")
	}

	consumer, err := u.consumerRepo.GetConsumerByID(transaction.ConsumerID)
	if err != nil {
		return errors.New("consumer not found")
	}

	if transaction.OTR > consumer.CreditLimit {
		return fmt.Errorf("transaction amount exceeds consumer's credit limit: %f > %f", transaction.OTR, consumer.CreditLimit)
	}

	consumer.CreditLimit -= transaction.OTR
	if err := u.consumerRepo.UpdateConsumer(&consumer); err != nil {
		return err
	}

	return u.repo.CreateTransaction(&transaction)
}