package models

import "time"

type Buying struct {
	Bid         uint      `gorm:"column:bid;primaryKey"`
	CustomerID  uint      `gorm:"column:uid"`
	BuyDateTime time.Time `gorm:"column:buyDateTime"`
	Image       string    `gorm:"column:image;size:3000"`

	//Courses  []Course `gorm:"foreignKey:bid"`
	Customer Customer `gorm:"foreignKey:CustomerID"`
}

func (Buying) TableName() string {
	return "Buying"
}
