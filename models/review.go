package models

type Review struct {
	Rid     uint   `gorm:"column:rid;primaryKey"`
	Uid     uint   `gorm:"column:uid"`
	CoID    uint   `gorm:"column:CoID"`
	Details string `gorm:"column:details;size:1000"`
	Score   uint   `gorm:"column:score"`
	Weight  uint   `gorm:"column:weight"`

	Coach    Coach    `gorm:"foreignKey:cid"`
	Customer Customer `gorm:"foreignKey:uid"`
	Course   Course   `gorm:"foreignKey:coID"`
}

func (Review) TableName() string {
	return "Review"
}
