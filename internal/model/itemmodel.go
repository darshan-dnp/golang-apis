package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	ID               string `gorm:"type:text;primaryKey" json:"id"`
	ReceiptID        string
	ShortDescription string `json:"shortDescription" validate:"required"`
	Price            string `json:"price" validate:"required"`
}

func (item *Item) BeforeCreate(tx *gorm.DB) (err error) {
	item.ID = uuid.New().String()
	return
}

func (i *Item) Validate() error {
	return validate.Struct(i)
}
