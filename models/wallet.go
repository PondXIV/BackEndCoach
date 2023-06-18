package models

type Wallet struct {
	Wid            uint    `gorm:"column:wid;primaryKey"`
	CustomerID     int     `gorm:"column:uid"`
	Money          float64 `gorm:"column:money"`
	Status         string  `gorm:"column:status;size:1"`
	Amount         float64 `gorm:"column:amount"`
	RetryFlag      string  `gorm:"column:retryFlag;size:20"`
	ReferenceNo    string  `gorm:"column:referenceNo;size:20"`
	GbpReferenceNo string  `gorm:"column:gbpReferenceNo;size:40"`
	CurrencyCode   string  `gorm:"column:CurrencyCode;size:20"`
	ResultCode     string  `gorm:"column:resultCode;size:20"`
	TotalAmount    float64 `gorm:"column:totalAmount"`
	Fee            float64 `gorm:"column:fee"`
	Vat            float64 `gorm:"column:vat"`
	ThbAmount      float64 `gorm:"column:thbAmount"`
	CustomerName   string  `gorm:"column:customerName;size:250"`
	Date           string  `gorm:"column:date;size:30"`
	Time           string  `gorm:"column:time;size:30"`
	PaymentType    string  `gorm:"column:paymentType;size:20"`

	Customer Customer `gorm:"foreignKey:CustomerID"`
}

func (Wallet) TableName() string {
	return "Wallet"
}
