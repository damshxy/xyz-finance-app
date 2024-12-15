package repository

import (
	"github.com/damshxy/xyz-finance-app/internal/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction *models.Transaction) error
	GetTransactionByConsumerID(consumerID int) ([]*models.Transaction, error)
	GetTransactionByID(id int) (*models.Transaction, error)
	MarkTransactionAsRefund(transactionID int) error
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

func (r *transactionRepository) GetTransactionByID(id int) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.First(&transaction, id).Error
	if err != nil {
		return &transaction, err
	}

	return &transaction, nil
}

func (r *transactionRepository) GetTransactionByConsumerID(consumerID int) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	err := r.db.Where("consumer_id = ?", consumerID).Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *transactionRepository) MarkTransactionAsRefund(transactionID int) error {
	err := r.db.Model(&models.Transaction{}).Where("id = ?", transactionID).Update("refunded", true).Error
	if err != nil {
		return err
	}

	return nil
}