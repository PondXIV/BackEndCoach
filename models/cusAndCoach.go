package models

type CoachAndCus struct {
	Cid uint `gorm:"column:cid;primaryKey"`
	Uid uint `gorm:"column:uid;primaryKey"`
}

func (CoachAndCus) TableName() string {
	return "Coach,Customer"
}
