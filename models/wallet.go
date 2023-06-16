package models

import "time"

type Wallet struct {
	Wid         uint      `gorm:"column:wid;primaryKey"`
	CustomerID  uint      `gorm:"column:uid"`
	Money       int       `gorm:"column:money"`
	Date        time.Time `gorm:"column:date"`
	Status      string    `gorm:"column:status;size:1"`
	ReferenceNo string    `gorm:"column:referenceNo;size:20"`

	Customer Customer `gorm:"foreignKey:CustomerID"`
}
