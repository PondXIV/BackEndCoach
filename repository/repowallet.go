package repository

import (
	"backEndGo/models"
	"fmt"

	"gorm.io/gorm"
)

type WalletRepository interface {
	UpdateWallet(ReferenceNo string, GbpRefNo string) (int64, error)
}
type WalletDB struct {
	db *gorm.DB
}

// UpdateWallet implements WalletRepository.
func (w WalletDB) UpdateWallet(ReferenceNo string, GbpRefNo string) (int64, error) {
	fmt.Println("GbpRefNo", GbpRefNo)
	result := w.db.Model(models.Wallet{}).Where("referenceNo = ?", ReferenceNo).Updates(
		models.Wallet{
			Status:         "1",
			GbpReferenceNo: ReferenceNo,
			//Customer: models.Customer{},
		})
	return result.RowsAffected, nil
}

func NewWalletRepository() WalletRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}

	return WalletDB{db}
}
