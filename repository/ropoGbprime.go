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
		ResultCode:     "",
		ReferenceNo:    "",
		GbpReferenceNo: "",
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
