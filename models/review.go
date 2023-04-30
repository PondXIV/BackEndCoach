package models

type Review struct {
	Rid        uint   `gorm:"column:rid;primaryKey"`
	CustomerID uint   `gorm:"column:uid"`
	CourseID   uint   `gorm:"column:coID"`
	Details    string `gorm:"column:details;size:1000"`
	Score      uint   `gorm:"column:score"`
	Weight     uint   `gorm:"column:weight"`

	//Coach    Coach    `gorm:"foreignKey:cid"`
	Customer Customer `gorm:"foreignKey:CustomerID"`
	Course   Course   `gorm:"foreignKey:CourseID"`
}

func (Review) TableName() string {
	return "Review"
}
