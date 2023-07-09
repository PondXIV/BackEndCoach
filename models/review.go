package models

type Review struct {
	Rid        uint   `gorm:"column:rid;primaryKey"`
	CustomerID int    `gorm:"column:uid"`
	CourseID   int    `gorm:"column:coID"`
	Details    string `gorm:"column:details;size:1000"`
	Score      int    `gorm:"column:score"`
	Weight     int    `gorm:"column:weight"`

	//Coach    Coach    `gorm:"foreignKey:cid"`
	Customer Customer `gorm:"foreignKey:CustomerID"`
	Course   Course   `gorm:"foreignKey:CourseID"`
}

func (Review) TableName() string {
	return "Review"
}
