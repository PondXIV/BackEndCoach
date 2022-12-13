package models

type Request struct {
	rpID    uint   `gorm:"column:rpID;primaryKey"`
	Cid     uint   `gorm:"column:cid"`
	Uid     uint   `gorm:"column:uid"`
	CpID    uint   `gorm:"column:cpID"`
	Status  string `gorm:"column:status;size:1"`
	Details string `gorm:"column:details;size:250"`

	Coach    Coach    `gorm:"foreignKey:cid"`
	Customer Customer `gorm:"foreignKey:uid"`
	Clip     Clip     `gorm:"foreignKey:cpID"`
}

func (Request) TableName() string {
	return "Request"
}
