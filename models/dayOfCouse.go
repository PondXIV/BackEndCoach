package models

type DayOfCouse struct {
	Did		uint      `gorm:"column:did;primaryKey"`
	CoID	uint 	  `gorm:"column:CoID"`
	Sequence	uint	  `gorm:"column:Sequence"`

	Course		Course	 	`gorm:"foreignKey:coID"`
	
	
	
}
func (DayOfCouse) TableName() string {
	return "DayOfCouse"
}