package models

import (
	"time"
)

//import "time"

type Coach struct {
	Cid           uint      `gorm:"column:cid;primaryKey"`
	Username      string    `gorm:"column:username;size:50"`
	Password      string    `gorm:"column:password;size:20"`
	Email         string    `gorm:"column:email;size:100"`
	FullName      string    `gorm:"column:fullName;size:100"`
	Birthday      time.Time `gorm:"column:birthday"`
	Gender        string    `gorm:"column:gender;size:1"`
	Phone         string    `gorm:"column:phone;size:10"`
	Image         string    `gorm:"column:image;size:3000"`
	Qualification string    `gorm:"column:qualification;size:1000"`
	Property      string    `gorm:"column:property;size:1000"`
	FacebookID    string    `gorm:"column:facebookId;size:100"`

	Chats []Chat `gorm:"foreignKey:cid"`
}

func (Coach) TableName() string {
	return "Coach"
}
