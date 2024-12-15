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

func TestCreateConsumer_Success(t *testing.T) {
	mockRepo := new(mocks.ConsumerRepository)
	mockConsumer := models.Consumer{
		NIK:         "1234567890",
		FullName:    "John Doe",
		CreditLimit: 0,
	}

	mockRepo.On("ConsumerExistByNIK", mockConsumer.NIK).Return(false, nil)
	mockRepo.On("CreateConsumer", mock.AnythingOfType("*models.Consumer")).Return(nil)

	consumerUsecase := usecase.NewConsumerUsecase(mockRepo)
	err := consumerUsecase.CreateConsumer(mockConsumer)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateConsumer_NIKAlreadyExist(t *testing.T) {
	mockRepo := new(mocks.ConsumerRepository)
	mockConsumer := models.Consumer{
		NIK:         "1234567890",
		FullName:    "John Doe",
	}

	mockRepo.On("ConsumerExistByNIK", mockConsumer.NIK).Return(true, nil)

	consumerUsecase := usecase.NewConsumerUsecase(mockRepo)
	err := consumerUsecase.CreateConsumer(mockConsumer)

	assert.Error(t, err)
	assert.Equal(t, "NIK already exists", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestGetConsumerByID_Success(t *testing.T) {
	mockRepo := new(mocks.ConsumerRepository)
	mockConsumer := models.Consumer{
		ID:          1,
		NIK:         "1234567890",
		FullName:    "John Doe",
	}

	mockRepo.On("GetConsumerByID", mockConsumer.ID).Return(mockConsumer, nil)

	consumerUsecase := usecase.NewConsumerUsecase(mockRepo)
	result, err := consumerUsecase.GetConsumerByID(mockConsumer.ID)

	assert.NoError(t, err)
	assert.Equal(t, mockConsumer, result)
	mockRepo.AssertExpectations(t)
}

func TestGetConsumerByID_ConsumerNotFound(t *testing.T) {
	mockRepo := new(mocks.ConsumerRepository)

	mockRepo.On("GetConsumerByID", 1).Return(models.Consumer{}, errors.New("consumer not found"))

	consumerUsecase := usecase.NewConsumerUsecase(mockRepo)
	result, err := consumerUsecase.GetConsumerByID(1)

	assert.Error(t, err)
	assert.Equal(t, "consumer not found", err.Error())
	assert.Equal(t, models.Consumer{}, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateConsumer_Success(t *testing.T) {
	mockRepo := new(mocks.ConsumerRepository)
	mockConsumer := models.Consumer{
		ID:          1,
		NIK:         "1234567890",
		FullName:    "John Doe",
		CreditLimit: 1000000.0,
	}

	updateCreditLimit := 2000000.0

	mockRepo.On("GetConsumerByID", mockConsumer.ID).Return(mockConsumer, nil)
	mockRepo.On("UpdateConsumer", mock.AnythingOfType("*models.Consumer")).Return(nil)

	consumerUsecase := usecase.NewConsumerUsecase(mockRepo)
	err := consumerUsecase.UpdateConsumer(mockConsumer.ID, updateCreditLimit)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateConsumer_ConsumerNotFound(t *testing.T) {
	mockRepo := new(mocks.ConsumerRepository)
	
	mockRepo.On("GetConsumerByID", 1).Return(models.Consumer{}, errors.New("consumer not found"))

	consumerUsecase := usecase.NewConsumerUsecase(mockRepo)
	err := consumerUsecase.UpdateConsumer(1, 2000000.0)

	assert.Error(t, err)
	assert.Equal(t, "consumer not found", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestUpdateConsumer_NegativeCreditLimit(t *testing.T) {
	mockRepo := new(mocks.ConsumerRepository)
	mockConsumer := models.Consumer{
		ID:          1,
		NIK:         "1234567890123456",
		FullName:    "John Doe",
		CreditLimit: 1000000,
	}

	mockRepo.On("GetConsumerByID", mockConsumer.ID).Return(mockConsumer, nil)

	consumerUsecase := usecase.NewConsumerUsecase(mockRepo)
	err := consumerUsecase.UpdateConsumer(mockConsumer.ID, -500000.0)

	assert.Error(t, err)
	assert.Equal(t, "credit limit can't be negative", err.Error())
	mockRepo.AssertExpectations(t)
}