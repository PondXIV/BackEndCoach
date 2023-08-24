package models

import "time"

type Buying struct {
	Bid         uint      `gorm:"column:bid;primaryKey"`
	CustomerID  uint      `gorm:"column:uid"`
	BuyDateTime time.Time `gorm:"column:buyDateTime"`
	CourseID    uint      `gorm:"column:coID;foreignKey"`
	OriginalID  uint      `gorm:"column:originalID"`
	Weight      int       `gorm:"column:weight"`

	Customer Customer `gorm:"foreignKey:CustomerID"`
	Course   Course   `gorm:"foreignKey:CourseID"`
	//Courses  []Course `gorm:"foreignKey:BuyingID"`
}

func (Buying) TableName() string {
	return "Buying"
}
