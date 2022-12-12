package models

import (
	"time"
)

//import "time"

type Customer struct {
	Uid        uint      `gorm:"column:uid;primaryKey"`
	AliasName  string    `gorm:"column:aliasName;size:50"`
	Password   string    `gorm:"column:password;size:20"`
	Email      string    `gorm:"column:email;size:100"`
	FullName   string    `gorm:"column:fullName;size:100"`
	Birthday   time.Time `gorm:"column:birthday"`
	Gender     string    `gorm:"column:gender;size:1"`
	Phone      string    `gorm:"column:phone;size:10"`
	Image      string    `gorm:"column:image;size:3000"`
	Weight     uint      `gorm:"column:weight"`
	Height     uint      `gorm:"column:height"`
	FacebookID string    `gorm:"column:facebookId;size:100"`
	Price      uint      `gorm:"column:price"`

	//Chats	[]Chat		 `gorm:"foreignKey:uid"`
	Buying  []Buying	 `gorm:"foreignKey:uid"`
}

func (Customer) TableName() string {
	return "Customer"
}
