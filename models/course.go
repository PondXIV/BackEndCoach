package models

import "time"

type Course struct {
	CoID           uint      `gorm:"column:coID;primaryKey"`
	CoachID        uint      `gorm:"column:cid"`
	BuyingID       uint      `gorm:"default:null;column:bid;foreignKey"`
	Name           string    `gorm:"column:name;size:50"`
	Details        string    `gorm:"column:details;size:250"`
	Level          string    `gorm:"column:level;size:1"`
	Amount         uint      `gorm:"column:amount"`
	Image          string    `gorm:"column:image;size:3000"`
	Days           int       `gorm:"column:days"`
	Price          uint      `gorm:"column:price"`
	Status         string    `gorm:"column:status;size:1"`
	ExpirationDate time.Time `gorm:"default:null;column:expiration date"`
	//DayOfCouses    []DayOfCouse `gorm:"foreignKey:coID"`

	//Coach  Coach  `gorm:"foreignKey:CoachID"`
	Buying Buying `gorm:"refernces:Bid"`

	DayOfCouses []DayOfCouse `gorm:"foreignKey:CourseID"`
}

func (Course) TableName() string {
	return "Course"
}
