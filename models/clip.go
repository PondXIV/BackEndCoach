package models

type Clip struct {
	CpID         uint   `gorm:"column:cpID;primaryKey"`
	ListClipID   uint   `gorm:"column:icpID"`
	DayOfCouseID uint   `gorm:"column:did"`
	Status       string `gorm:"column:status;size:1"`

	ListClip ListClip `gorm:"foreignKey:ListClipID"`
	// DayOfCouse DayOfCouse `gorm:"foreignKey:DayOfCouseID"`
}

func (Clip) TableName() string {
	return "Clip"
}
