package entity

type User struct {
	BaseEntity
	Name              string `gorm:"column:name;default:NULL"`
	RankMasterAccount string `gorm:"column:rank_master_account;default:NULL"`
	Password          string `gorm:"column:password;default:NULL"`
	Avatar            string `gorm:"column:avatar;default:NULL"`
	Mobile            string `gorm:"column:mobile;default:NULL"`
	CryptSalt         string `gorm:"column:crypt_salt;default:NULL"`
}

func (u *User) TableName() string {
	return "user"
}
