package models

import "time"

type Buying struct {
	Bid         uint      `gorm:"column:bid;primaryKey"`
	CustomerID  uint      `gorm:"column:uid"`
	BuyDateTime time.Time `gorm:"column:buyDateTime"`

	Customer Customer `gorm:"foreignKey:CustomerID"`
	//Courses  []Course `gorm:"foreignKey:BuyingID"`
}

func (Buying) TableName() string {
	return "Buying"
}
