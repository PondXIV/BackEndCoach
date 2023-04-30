package models

type DayOfCouse struct {
	Did      uint `gorm:"column:did;primaryKey"`
	CourseID uint `gorm:"column:coID"`
	Sequence int  `gorm:"column:sequence"`

	Course Course `gorm:"foreignKey:CourseID"`
}

func (DayOfCouse) TableName() string {
	return "DayOfCouse"
}
