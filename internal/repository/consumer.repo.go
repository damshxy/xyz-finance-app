package repository

import (
	"github.com/damshxy/xyz-finance-app/internal/models"
	"gorm.io/gorm"
)

type ConsumerRepository interface {
	CreateConsumer(consumer *models.Consumer) error
	ConsumerExistByNIK(nik string) (bool, error)
	GetConsumerByID(id int) (models.Consumer, error)
	UpdateConsumer(consumer *models.Consumer) error
}

type consumerRepository struct {
	db *gorm.DB
}

func NewConsumerRepository(db *gorm.DB) ConsumerRepository {
	return &consumerRepository{
		db: db,
	}
}

func (r *consumerRepository) CreateConsumer(consumer *models.Consumer) error {
	err := r.db.Create(consumer).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *consumerRepository) ConsumerExistByNIK(nik string) (bool, error) {
	var count int64
	err := r.db.Model(&models.Consumer{}).Where("nik = ?", nik).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *consumerRepository) GetConsumerByID(id int) (models.Consumer, error) {
	var consumer models.Consumer
	err := r.db.First(&consumer, id).Error
	if err != nil {
		return consumer, err
	}

	return consumer, nil
}

func (r *consumerRepository) UpdateConsumer(consumer *models.Consumer) error {
	err := r.db.Save(consumer).Error
	if err != nil {
		return err
	}

	return nil
}