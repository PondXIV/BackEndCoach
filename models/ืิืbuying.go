package models

import "time"

type Buying struct {
	Bid         uint      `gorm:"column:bid;primaryKey"`
	CustomerID  uint      `gorm:"column:uid"`
	BuyDateTime time.Time `gorm:"column:buyDateTime"`
	Image       string    `gorm:"column:image;size:3000"`

	Customer Customer `gorm:"foreignKey:CustomerID"`
	//Courses  []Course `gorm:"foreignKey:BuyingID"`
}

func (Buying) TableName() string {
	return "Buying"
}
