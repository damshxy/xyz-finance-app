package usecase_test

import (
	"errors"
	"testing"

	"github.com/damshxy/xyz-finance-app/internal/models"
	"github.com/damshxy/xyz-finance-app/internal/repository/mocks"
	"github.com/damshxy/xyz-finance-app/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransaction_Success(t *testing.T) {
	mockTransactionRepo := new(mocks.TransactionRepository)
	mockConsumerRepo := new(mocks.ConsumerRepository)

	mockConsumer := models.Consumer{
		ID:          1,
		NIK:         "1234567890",
		FullName:    "John Doe",
		CreditLimit: 5000000,
	}

	mockTransaction := models.Transaction{
		ConsumerID:     mockConsumer.ID,
		ContractNumber: "CN123456",
		OTR:            2000000,
		AdminFee:       50000,
		Installment:    250000,
		Interest:       5,
		AssetName:      "Honda Brio",
	}

	mockConsumerRepo.On("GetConsumerByID", mockConsumer.ID).Return(mockConsumer, nil)
	mockConsumerRepo.On("UpdateConsumer", mock.AnythingOfType("*models.Consumer")).Return(nil)
	mockTransactionRepo.On("CreateTransaction", mock.AnythingOfType("*models.Transaction")).Return(nil)

	transactionUsecase := usecase.NewTransactionUsecase(mockTransactionRepo, mockConsumerRepo)
	err := transactionUsecase.CreateTransaction(mockTransaction)

	assert.NoError(t, err)
	mockConsumerRepo.AssertExpectations(t)
	mockTransactionRepo.AssertExpectations(t)
}

func TestCreateTransaction_ExceedsCreditLimit(t *testing.T) {
	mockTransactionRepo := new(mocks.TransactionRepository)
	mockConsumerRepo := new(mocks.ConsumerRepository)

	mockConsumer := models.Consumer{
		ID:          1,
		NIK:         "1234567890",
		FullName:    "John Doe",
		CreditLimit: 1000000, // Credit limit is insufficient
	}

	mockTransaction := models.Transaction{
		ConsumerID:     mockConsumer.ID,
		ContractNumber: "CN123456",
		OTR:            2000000, // OTR exceeds credit limit
	}

	mockConsumerRepo.On("GetConsumerByID", mockConsumer.ID).Return(mockConsumer, nil)

	transactionUsecase := usecase.NewTransactionUsecase(mockTransactionRepo, mockConsumerRepo)
	err := transactionUsecase.CreateTransaction(mockTransaction)

	assert.Error(t, err)
	assert.Equal(t, "transaction amount exceeds consumer's credit limit: 2000000.000000 > 1000000.000000", err.Error())
	mockConsumerRepo.AssertExpectations(t)
	mockTransactionRepo.AssertNotCalled(t, "CreateTransaction", mock.Anything)
}

func TestRefundTransaction_Success(t *testing.T) {
	mockTransactionRepo := new(mocks.TransactionRepository)
	mockConsumerRepo := new(mocks.ConsumerRepository)

	mockTransaction := &models.Transaction{
		ID:          1,
		ConsumerID:  1,
		OTR:         2000000,
		Refunded:    false,
	}

	mockConsumer := models.Consumer{
		ID:          mockTransaction.ConsumerID,
		CreditLimit: 3000000,
	}

	mockTransactionRepo.On("GetTransactionByID", mockTransaction.ID).Return(mockTransaction, nil)
	mockTransactionRepo.On("MarkTransactionAsRefund", mockTransaction.ID).Return(nil)

	mockConsumerRepo.On("GetConsumerByID", mockTransaction.ConsumerID).Return(mockConsumer, nil)
	mockConsumerRepo.On("UpdateConsumer", mock.AnythingOfType("*models.Consumer")).Return(nil)

	transactionUsecase := usecase.NewTransactionUsecase(mockTransactionRepo, mockConsumerRepo)
	err := transactionUsecase.RefundTransaction(mockTransaction.ID)

	assert.NoError(t, err)
	mockTransactionRepo.AssertExpectations(t)
	mockConsumerRepo.AssertExpectations(t)
}


func TestRefundTransaction_NotFound(t *testing.T) {
	mockTransactionRepo := new(mocks.TransactionRepository)
	mockConsumerRepo := new(mocks.ConsumerRepository)

	mockTransactionRepo.On("GetTransactionByID", 1).Return(nil, errors.New("transaction not found"))

	transactionUsecase := usecase.NewTransactionUsecase(mockTransactionRepo, mockConsumerRepo)
	err := transactionUsecase.RefundTransaction(1)

	assert.Error(t, err)
	assert.Equal(t, "transaction not found", err.Error())
	mockTransactionRepo.AssertExpectations(t)
	mockConsumerRepo.AssertNotCalled(t, "UpdateConsumer", mock.Anything)
}



func TestGetTransactionByConsumerID_Success(t *testing.T) {
	mockTransactionRepo := new(mocks.TransactionRepository)
	mockConsumerRepo := new(mocks.ConsumerRepository)

	mockConsumer := models.Consumer{
		ID:          1,
		NIK:         "1234567890",
		FullName:    "John Doe",
		CreditLimit: 5000000,
	}

	mockTransactions := []*models.Transaction{
		{
			ConsumerID:     mockConsumer.ID,
			ContractNumber: "CN123456",
			OTR:            2000000,
			AdminFee:       50000,
			Installment:    250000,
			Interest:       5,
			AssetName:      "Honda Brio",
		},
		{
			ConsumerID:     mockConsumer.ID,
			ContractNumber: "CN123457",
			OTR:            3000000,
			AdminFee:       60000,
			Installment:    280000,
			Interest:       5,
			AssetName:      "Toyota Corolla",
		},
	}

	mockConsumerRepo.On("GetConsumerByID", mockConsumer.ID).Return(mockConsumer, nil)
	mockTransactionRepo.On("GetTransactionByConsumerID", mockConsumer.ID).Return(mockTransactions, nil)

	transactionUsecase := usecase.NewTransactionUsecase(mockTransactionRepo, mockConsumerRepo)
	transactions, err := transactionUsecase.GetTransactionByConsumerID(mockConsumer.ID)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(transactions))
	mockConsumerRepo.AssertExpectations(t)
	mockTransactionRepo.AssertExpectations(t)
}
