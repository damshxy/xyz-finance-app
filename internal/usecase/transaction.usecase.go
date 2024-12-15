package usecase

import (
	"errors"
	"fmt"

	"github.com/damshxy/xyz-finance-app/internal/models"
	"github.com/damshxy/xyz-finance-app/internal/repository"
)

type TransactionUsecase interface {
	CreateTransaction(transaction models.Transaction) error
	GetTransactionByConsumerID(consumerID int) ([]*models.Transaction, error)
	RefundTransaction(transactionID int) error
}

type transactionUsecase struct {
	repo          repository.TransactionRepository
	consumerRepo  repository.ConsumerRepository
}

func NewTransactionUsecase(repo repository.TransactionRepository, consumerRepo repository.ConsumerRepository) TransactionUsecase {
	return &transactionUsecase{
		repo:         repo,
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

func (u *transactionUsecase) RefundTransaction(transactionID int) error {
	transaction, err := u.repo.GetTransactionByID(transactionID)
	if err != nil {
		return errors.New("transaction not found")
	}

	consumer, err := u.consumerRepo.GetConsumerByID(transaction.ConsumerID)
	if err != nil {
		return errors.New("consumer not found")
	}

	consumer.CreditLimit += transaction.OTR
	if err := u.consumerRepo.UpdateConsumer(&consumer); err != nil {
		return fmt.Errorf("failed to update consumer credit limit: %w", err)
	}

	if err := u.repo.MarkTransactionAsRefund(transactionID); err != nil {
		return fmt.Errorf("failed to mark transaction as refunded: %w", err)
	}

	return nil
}

func (u *transactionUsecase) GetTransactionByConsumerID(consumerID int) ([]*models.Transaction, error) {
	_, err := u.consumerRepo.GetConsumerByID(consumerID)
	if err != nil {
		return nil, errors.New("consumer not found")
	}

	transaction, err := u.repo.GetTransactionByConsumerID(consumerID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve transactions: %w", err)
	}

	return transaction, nil
}