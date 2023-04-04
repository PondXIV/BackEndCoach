package models

type ListFood struct {
	Ifid     uint   `gorm:"column:ifid;primaryKey"`
	Cid      int    `gorm:"column:cid"`
	Name     string `gorm:"column:name;size:100"`
	Image    string `gorm:"column:image;size:3000"`
	Details  string `gorm:"column:details;size:1000"`
	Calories int    `gorm:"column:calorie"`

	//Coach Coach `gorm:"foreignKey:cid"`
}

func (ListFood) TableName() string {
	return "listFood"
}
