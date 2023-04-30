package models

type DayOfCouse struct {
	Did      uint `gorm:"column:did;primaryKey"`
	CourseID uint `gorm:"column:coID"`
	Sequence int  `gorm:"column:sequence"`

	Foods []Food `gorm:"foreignKey:DayOfCouseID"`
	Clips []Clip `gorm:"foreignKey:DayOfCouseID"`
	//Course Course `gorm:"foreignKey:CourseID"`
}

func (DayOfCouse) TableName() string {
	return "DayOfCouse"
}
