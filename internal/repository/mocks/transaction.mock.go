package mocks

import (
	"github.com/damshxy/xyz-finance-app/internal/models"
	"github.com/stretchr/testify/mock"
)

type TransactionRepository struct {
	mock.Mock
}

func (m *TransactionRepository) GetTransactionByID(transactionID int) (*models.Transaction, error) {
	args := m.Called(transactionID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Transaction), args.Error(1)
}

func (m *TransactionRepository) GetTransactionByConsumerID(consumerID int) ([]*models.Transaction, error) {
	args := m.Called(consumerID)
	return args.Get(0).([]*models.Transaction), args.Error(1)
}

func (m *TransactionRepository) CreateTransaction(transaction *models.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *TransactionRepository) MarkTransactionAsRefund(transactionID int) error {
	args := m.Called(transactionID)
	return args.Error(0)
}
