package models

import "time"

type Course struct {
	CoID           uint      `gorm:"column:coID;primaryKey"`
	Cid            uint      `gorm:"column:cid"`
	Bid            uint      `gorm:"column:bid"`
	Name           string    `gorm:"column:name;size:50"`
	Details        string    `gorm:"column:details;size:250"`
	Level          string    `gorm:"column:level;size:1"`
	Amount         uint      `gorm:"column:amount"`
	Image          string    `gorm:"column:image;size:3000"`
	Days           uint      `gorm:"column:days"`
	Price          uint      `gorm:"column:price"`
	Status         string    `gorm:"column:status;size:1"`
	ExpirationDate time.Time `gorm:"column:expiration date"`

	//Coach  Coach  `gorm:"foreignKey:cid"`
	//Buying Buying `gorm:"foreignKey:bid"`
}

func (Course) TableName() string {
	return "Course"
}
