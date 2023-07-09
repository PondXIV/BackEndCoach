package models

type Request struct {
	RpID       uint   `gorm:"column:rqID;primaryKey"`
	CoachID    int    `gorm:"column:cid"`
	CustomerID int    `gorm:"column:uid"`
	ClipID     int    `gorm:"column:cpID"`
	Status     string `gorm:"column:status;size:1"`
	Details    string `gorm:"column:details;size:250"`

	//Coach    Coach    `gorm:"foreignKey:CoachID"`
	Customer Customer `gorm:"foreignKey:CustomerID"`
	Clip     Clip     `gorm:"foreignKey:ClipID"`
}

func (Request) TableName() string {
	return "Request"
}
