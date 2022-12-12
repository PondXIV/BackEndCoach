package models

type IClip struct {
	IcpID	uint      `gorm:"column:icpID;primaryKey"`
	Cid		uint 	  `gorm:"column:cid"`
	Name  	string    `gorm:"column:name;size:100"`
	AmountPerSet 	string 	  `gorm:"column:amount per set;size:50"`
	Video	string    `gorm:"column:video;size:3000"`
	Details string    `gorm:"column:details;size:2000"`
	
	Coach		Coach	 	`gorm:"foreignKey:cid"`
	
}
func (IClip) TableName() string {
	return "iClip"
}