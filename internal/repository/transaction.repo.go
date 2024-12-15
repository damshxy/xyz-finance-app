package repository

import (
	"github.com/damshxy/xyz-finance-app/internal/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction *models.Transaction) error
	GetTransactionByConsumerID(consumerID int) ([]*models.Transaction, error)
	UpdateTransaction(transaction *models.Transaction) error
	DeleteTransaction(id int) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) CreateTransaction(transaction *models.Transaction) error {
	err := r.db.Create(transaction).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *transactionRepository) GetTransactionByConsumerID(consumerID int) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	err := r.db.Where("consumer_id = ?", consumerID).Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *transactionRepository) UpdateTransaction(transaction *models.Transaction) error {
	err := r.db.Save(transaction).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *transactionRepository) DeleteTransaction(id int) error {
	err := r.db.Delete(&models.Transaction{}, id).Error
	if err != nil {
		return err
	}

	return nil
}