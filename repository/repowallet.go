package repository

import (
	"backEndGo/models"

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

	result := w.db.Model(models.Wallet{}).Where("referenceNo = ?", ReferenceNo).Updates(
		models.Wallet{
			Status:         "1",
			GbpReferenceNo: GbpRefNo,
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
