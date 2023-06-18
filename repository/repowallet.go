package repository

import (
	"backEndGo/models"

	"gorm.io/gorm"
)

type WalletRepository interface {
	UpdateWallet(ReferenceNo string, ResGb *models.Gbprimpay) (int64, error)
}
type WalletDB struct {
	db *gorm.DB
}

// UpdateWallet implements WalletRepository.
func (w WalletDB) UpdateWallet(ReferenceNo string, ResGb *models.Gbprimpay) (int64, error) {

	result := w.db.Model(models.Wallet{}).Where("referenceNo = ?", ReferenceNo).Updates(
		models.Wallet{
			Status:         "1",
			GbpReferenceNo: ResGb.GbpReferenceNo,
			Amount:         int(ResGb.Amount),
			RetryFlag:      ResGb.RetryFlag,
			CurrencyCode:   ResGb.CurrencyCode,
			TotalAmount:    int(ResGb.TotalAmount),
			Fee:            ResGb.Fee,
			Vat:            ResGb.Vat,
			ThbAmount:      int(ResGb.ThbAmount),
			CustomerName:   ResGb.CustomerName,
			Date:           ResGb.Date,
			Time:           ResGb.Time,
			PaymentType:    ResGb.PaymentType,
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
