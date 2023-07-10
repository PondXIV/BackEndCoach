package models

type Chat struct {
	ChID       uint   `gorm:"column:chID;primaryKey"`
	CustomerID uint   `gorm:"column:uid"`
	BuyingID   uint   `gorm:"column:bid"`
	CoachID    uint   `gorm:"column:cid"`
	Message    string `gorm:"column:message;size250"`

	// Customer Customer `gorm:"foreignKey:CustomerID"`
	Buying Buying `gorm:"foreignKey:BuyingID"`
	// Coach    Coach    `gorm:"foreignKey:CoachID"`
}

func (Chat) TableName() string {
	return "Chat"
}
