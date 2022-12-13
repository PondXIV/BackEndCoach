package models

import "time"

type Buying struct {
	Bid         uint      `gorm:"column:bid;primaryKey"`
	Uid         uint      `gorm:"column:uid"`
	BuyDateTime time.Time `gorm:"column:buyDateTime"`
	Image       string    `gorm:"column:image;size:3000"`

	Customer Customer `gorm:"foreignKey:uid"`
}

func (Buying) TableName() string {
	return "Buying"
}
