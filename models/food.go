package models

type Food struct {
	Fid          uint   `gorm:"column:fid;primaryKey"`
	ListFoodID   int    `gorm:"column:ifid"`
	DayOfCouseID uint   `gorm:"column:did"`
	Time         string `gorm:"column:time;size:15"`

	ListFood ListFood `gorm:"foreignKey:ListFoodID"`
	//DayOfCouse DayOfCouse `gorm:"foreignKey:DayOfCouseID"`
}

func (Food) TableName() string {
	return "Food"
}
