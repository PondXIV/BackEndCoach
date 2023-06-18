package repository

import (
	"backEndGo/models"
	"fmt"

	"gorm.io/gorm"
)

type GbprimeRepository interface {
	Gbprimepay(*models.Gbprimpay)
}
type GbprimeDB struct {
	db *gorm.DB
}

// Gbprimepay implements GbprimeRepository.
func (g GbprimeDB) Gbprimepay(*models.Gbprimpay) {

	resGb := *&models.Gbprimpay{
		Amount:         0,
		RetryFlag:      "",
		ReferenceNo:    "",
		GbpReferenceNo: "",
		CurrencyCode:   "",
		ResultCode:     "",
		TotalAmount:    0,
		Fee:            0,
		Vat:            0,
		ThbAmount:      0,
		CustomerName:   "",
		Date:           "",
		Time:           "",
		PaymentType:    "",
	}
	fmt.Println(resGb)

}

func NewGbprimeRepository() GbprimeRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}

	return GbprimeDB{db}
}
