package entity

import (
	"time"

	"gorm.io/gorm"
)

type BaseEntity struct {
	Id        string         `gorm:"column:id;primary_key;NOT NULL"`
	CreatedAt time.Time      `gorm:"column:created_at;"`
	UpdatedAt time.Time      `gorm:"column:updated_at;"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}
