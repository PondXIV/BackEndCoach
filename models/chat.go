package models

type Chat struct {
	ChID    uint   `gorm:"column:chID;primaryKey"`
	Uid     uint   `gorm:"column:uid"`
	Bid     uint   `gorm:"column:bid"`
	Cid     uint   `gorm:"column:cid"`
	Message string `gorm:"column:message;size250"`

	Customer Customer `gorm:"foreignKey:uid"`
	Buying   Buying   `gorm:"foreignKey:bid"`
	Coach    Coach    `gorm:"foreignKey:cid"`
}

func (Chat) TableName() string {
	return "Chat"
}
