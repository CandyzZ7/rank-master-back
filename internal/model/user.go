package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id                string         `gorm:"column:id;primary_key;NOT NULL"`
	CreatedAt         time.Time      `gorm:"column:created_at;"`
	UpdatedAt         time.Time      `gorm:"column:updated_at;"`
	DeletedAt         gorm.DeletedAt `gorm:"column:deleted_at;index"`
	Name              string         `gorm:"column:name;default:NULL"`
	RankMasterAccount string         `gorm:"column:rank_master_account;default:NULL"`
	Password          string         `gorm:"column:password;default:NULL"`
	Avatar            string         `gorm:"column:avatar;default:NULL"`
	Mobile            string         `gorm:"column:mobile;default:NULL"`
	CryptSalt         string         `gorm:"column:crypt_salt;default:NULL"`
}

func (u *User) TableName() string {
	return "user"
}
