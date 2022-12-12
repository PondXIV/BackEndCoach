package models

type Clip struct {
	CpID	uint      `gorm:"column:cpID;primaryKey"`
	IcpID	uint 	  `gorm:"column:icpID"`
	Did		uint	  `gorm:"column:did"`
	Status	uint	  `gorm:"column:status"`

	
	IClip		IClip 		`gorm:"foreignKey:icpID"`
	DayOfCouse	DayOfCouse 	`gorm:"foreignKey:did"`
	
}
func (Clip) TableName() string {
	return "Clip"
}