package mocks

import (
	"github.com/damshxy/xyz-finance-app/internal/models"
	"github.com/stretchr/testify/mock"
)

type ConsumerRepository struct {
	mock.Mock
}

func (m *ConsumerRepository) ConsumerExistByNIK(nik string) (bool, error) {
	args := m.Called(nik)
	return args.Bool(0), args.Error(1)
}

func (m *ConsumerRepository) CreateConsumer(consumer *models.Consumer) error {
	args := m.Called(consumer)
	return args.Error(0)
}

func (m *ConsumerRepository) GetConsumerByID(id int) (models.Consumer, error) {
	args := m.Called(id)
	return args.Get(0).(models.Consumer), args.Error(1)
}

func (m *ConsumerRepository) UpdateConsumer(consumer *models.Consumer) error {
	args := m.Called(consumer)
	return args.Error(0)
}
