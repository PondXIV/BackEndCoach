package models

type Clip struct {
	CpID         uint `gorm:"column:cpID;primaryKey"`
	ListClipID   uint `gorm:"column:icpID"`
	DayOfCouseID uint `gorm:"column:did"`
	Status       uint `gorm:"column:status"`

	ListClip ListClip `gorm:"foreignKey:ListClipID"`
	// DayOfCouse DayOfCouse `gorm:"foreignKey:DayOfCouseID"`
}

func (Clip) TableName() string {
	return "Clip"
}
