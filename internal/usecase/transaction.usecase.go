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

// CreateTransaction processes the creation of a new transaction
func (u *transactionUsecase) CreateTransaction(transaction models.Transaction) error {
	// Validate transaction data
	if transaction.ConsumerID == 0 || transaction.OTR <= 0 {
		return errors.New("invalid transaction data")
	}

	// Get consumer data
	consumer, err := u.consumerRepo.GetConsumerByID(transaction.ConsumerID)
	if err != nil {
		return errors.New("consumer not found")
	}

	// Check if the transaction exceeds the consumer's credit limit
	if transaction.OTR > consumer.CreditLimit {
		return fmt.Errorf("transaction amount exceeds consumer's credit limit: %f > %f", transaction.OTR, consumer.CreditLimit)
	}

	// Deduct from consumer's credit limit
	consumer.CreditLimit -= transaction.OTR
	if err := u.consumerRepo.UpdateConsumer(&consumer); err != nil {
		return err
	}

	// Create the transaction record
	return u.repo.CreateTransaction(&transaction)
}

// RefundTransaction processes the refund of a transaction and restores credit limit
func (u *transactionUsecase) RefundTransaction(transactionID int) error {
	// Get the transaction by its ID
	transaction, err := u.repo.GetTransactionByID(transactionID)
	if err != nil {
		return errors.New("transaction not found")
	}

	// Get consumer data
	consumer, err := u.consumerRepo.GetConsumerByID(transaction.ConsumerID)
	if err != nil {
		return errors.New("consumer not found")
	}

	// Restore consumer's credit limit
	consumer.CreditLimit += transaction.OTR
	if err := u.consumerRepo.UpdateConsumer(&consumer); err != nil {
		return fmt.Errorf("failed to update consumer credit limit: %w", err)
	}

	// Mark the transaction as refunded
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