package models

type ListClip struct {
	IcpID        uint   `gorm:"column:icpID;primaryKey"`
	CoachID      uint   `gorm:"column:cid"`
	Name         string `gorm:"column:name;size:100"`
	AmountPerSet string `gorm:"column:amount_set;size:50"`
	Video        string `gorm:"column:video;size:3000"`
	Details      string `gorm:"column:details;size:2000"`

	// Coach Coach `gorm:"foreignKey:CoachID"`
}

func (ListClip) TableName() string {
	return "listClip"
}
