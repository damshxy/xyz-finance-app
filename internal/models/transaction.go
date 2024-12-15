package models

type Transaction struct {
	ID             int     `json:"id" gorm:"primaryKey"`
	ConsumerID     int     `json:"consumer_id" gorm:"not null"`
	Consumer 		Consumer `json:"consumer" gorm:"foreignKey:ConsumerID"`
	ContractNumber string  `json:"contract_number" gorm:"not null"`
	OTR            float64 `json:"otr" gorm:"not null"`
	AdminFee       float64 `json:"admin_fee" gorm:"not null"`
	Installment    float64 `json:"installment" gorm:"not null"`
	Interest       float64 `json:"interest" gorm:"not null"`
	AssetName      string  `json:"asset_name" gorm:"not null"`
}