package entity

type Template struct {
	BaseEntity
	Function string `gorm:"column:function;default:NULL"`
	Type     string `gorm:"column:type;default:NULL"`
	Topic    string `gorm:"column:topic;default:NULL"`
	Content  string `gorm:"column:content;default:NULL"`
	Version  int64  `gorm:"column:version;default:NULL"`
	Remark   string `gorm:"column:remark;default:NULL"`
}

func (t *Template) TableName() string {
	return "template"
}
