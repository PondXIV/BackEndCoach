package repository

import (
	"backEndGo/models"
	"fmt"

	"gorm.io/gorm"
)

type WalletRepository interface {
	UpdateWallet(ReferenceNo string, ResGb *models.Gbprimpay) (int64, error)
	InsertWallet(CusID int, wallet *models.Wallet) (int64, error)
	UpdateWalletUid(CusID int, price float64) (int64, error)
	GetUser(ReferenceNo string) (*models.Wallet, error)
}
type WalletDB struct {
	db *gorm.DB
}

// getUser implements WalletRepository.
func (w WalletDB) GetUser(ReferenceNo string) (*models.Wallet, error) {
	wallet := models.Wallet{}
	result := w.db.Where("referenceNo = ?", ReferenceNo).Find(&wallet)
	fmt.Println(result, "///", ReferenceNo)
	return &wallet, nil
}

// UpdateWalletUid implements WalletRepository.
func (w WalletDB) UpdateWalletUid(CusID int, price float64) (int64, error) {
	//mdcus := *&models.Customer{}
	result := w.db.Model(models.Customer{}).Where("uid = ?", CusID).Update("price", price)
	if result.Error != nil {
		panic(result.Error)
	}
	if result.RowsAffected > 0 {
		fmt.Println("CusID", CusID)
		fmt.Println("Update completed")
	}
	if result.RowsAffected == 0 {
		fmt.Println("Unupdate completed")
	}
	return result.RowsAffected, nil
}

// InsertWallet implements WalletRepository.
func (w WalletDB) InsertWallet(CusID int, wallet *models.Wallet) (int64, error) {
	result := w.db.Create(&models.Wallet{
		Wid:            0,
		CustomerID:     CusID,
		Money:          wallet.Money,
		Status:         "0",
		Amount:         0,
		RetryFlag:      "",
		ReferenceNo:    wallet.ReferenceNo,
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
	})
	if result.Error != nil {
		return -1, result.Error
	}
	return result.RowsAffected, nil
}

// UpdateWallet implements WalletRepository.
func (w WalletDB) UpdateWallet(ReferenceNo string, ResGb *models.Gbprimpay) (int64, error) {

	result := w.db.Model(models.Wallet{}).Where("referenceNo = ?", ReferenceNo).Updates(
		models.Wallet{
			Status:         "1",
			GbpReferenceNo: ResGb.GbpReferenceNo,
			Amount:         ResGb.Amount,
			RetryFlag:      ResGb.RetryFlag,
			CurrencyCode:   ResGb.CurrencyCode,
			ResultCode:     ResGb.ResultCode,
			TotalAmount:    ResGb.TotalAmount,
			Fee:            ResGb.Fee,
			Vat:            ResGb.Vat,
			ThbAmount:      ResGb.ThbAmount,
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
