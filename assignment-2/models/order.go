package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID           uint           `gorm:"primarykey"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	CustomerName string         `gorm:"type:varchar(100)" json:"customerName"`
	OrderedAt    time.Time      `json:"orderedAt"`
	Items        []Item         `json:"items"`
}
