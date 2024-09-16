package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Receipt struct {
	ID           string `gorm:"type:text;primaryKey" json:"id"`
	Retailer     string `json:"retailer" validate:"required"`
	PurchaseDate string `json:"purchaseDate" validate:"required"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `gorm:"foreignKey:ReceiptID" json:"items" validate:"required"`
	Total        string `json:"total" validate:"required"`
	Points       int    `json:"points"`
	DataHash     string `json:"-"`
}

func (receipt *Receipt) BeforeCreate(tx *gorm.DB) (err error) {
	receipt.ID = uuid.New().String()
	return
}

func (receipt *Receipt) Validate() error {
	return validate.Struct(receipt)
}
