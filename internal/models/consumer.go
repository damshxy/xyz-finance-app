package models

type Consumer struct {
	ID          int     `json:"id" gorm:"primaryKey"`
	NIK         string  `json:"nik" gorm:"unique;not null"`
	FullName    string  `json:"full_name" gorm:"not null"`
	LegalName   string  `json:"legal_name" gorm:"not null"`
	BirthPlace  string  `json:"birth_place" gorm:"not null"`
	BirthDate   string  `json:"birth_date" gorm:"not null"`
	Salary      float64 `json:"salary" gorm:"not null"`
	KTPPhoto    string  `json:"ktp_photo" gorm:"not null"`
	SelfiePhoto string  `json:"selfie_photo" gorm:"not null"`
	CreditLimit float64 `json:"credit_limit" gorm:"not null"`
}