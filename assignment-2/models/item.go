package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID          uint           `gorm:"primarykey" json:"lineItemId"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	ItemCode    string         `gorm:"type:varchar(100)" json:"itemCode"`
	Description string         `json:"description"`
	Quantity    int            `gorm:"type:integer" json:"quantity"`
	OrderID     uint           `json:"-"`
}
