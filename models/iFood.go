package models

type IFood struct {
	Ifid     uint   `gorm:"column:ifid;primaryKey"`
	Cid      uint   `gorm:"column:cid"`
	Name     string `gorm:"column:name;size:100"`
	Image    string `gorm:"column:image;size:3000"`
	Details  string `gorm:"column:details;size:1000"`
	Calories int    `gorm:"column:calories"`

	Coach Coach `gorm:"foreignKey:cid"`
}

func (IFood) TableName() string {
	return "iFood"
}
