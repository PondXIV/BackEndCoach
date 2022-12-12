package models

type Food struct {
	Fid  uint   `gorm:"column:fid;primaryKey"`
	IFid uint   `gorm:"column:ifid"`
	Did  uint   `gorm:"column:did"`
	Time string `gorm:"column:time;size:15"`

	IClip      IClip      `gorm:"foreignKey:icpID"`
	DayOfCouse DayOfCouse `gorm:"foreignKey:did"`
}

func (Food) TableName() string {
	return "Food"
}
