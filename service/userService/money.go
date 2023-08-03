package userservice

import (
	"backEndGo/models"
	"backEndGo/repository"
	"fmt"
)

var repoWallet = repository.NewWalletRepository()

type GbprimeService interface {
	ServiceGbprime(Gbprime *models.Gbprimpay)
	ServiceWallet(ReferenceNo string, ResGb *models.Gbprimpay) (int64, error)
	ServiceInsertWallet(CusID int, wallet *models.Wallet) (int64, error)
	ServiceHistoryWallet(CusID int) (*[]models.Wallet, error)
	//ServiceInsertWalletUid(CusID int, price int) (int64, error)
}

type GbprimeData struct {
}

// ServiceHistoryWallet implements GbprimeService.
func (GbprimeData) ServiceHistoryWallet(CusID int) (*[]models.Wallet, error) {
	wallet, err := repoWallet.GetHistoryWallet(CusID)
	fmt.Println(CusID)
	if err != nil {
		panic(err)
	}
	return wallet, nil
}

// ServiceInsertWallet implements GbprimeService.
func (g GbprimeData) ServiceInsertWallet(CusID int, wallet *models.Wallet) (int64, error) {
	rowsAffecteds, err := repoWallet.InsertWallet(CusID, wallet)
	if err != nil {
		panic(err)
	}
	return int64(rowsAffecteds), nil
}

// ServiceWallet implements GbprimeService.
func (g GbprimeData) ServiceWallet(ReferenceNo string, ResGb *models.Gbprimpay) (int64, error) {
	repoCus := repository.NewCustomerRepository()
	RowsAffected, err := repoWallet.UpdateWallet(ReferenceNo, ResGb)
	if err != nil {
		return -1, err
	}
	if RowsAffected > 0 {
		modelwallet, errs := repoWallet.GetUser(ReferenceNo)
		if errs != nil {
			return -1, err
		}
		if modelwallet.CustomerID != 0 {
			User, _ := repoCus.GetCustomerByID(modelwallet.CustomerID)

			sum := User.Price + (modelwallet.Money * 1000.00)
			rowsAffected, errs := repoWallet.UpdateWalletUid(int(User.Uid), sum)
			fmt.Println("CustomerID", modelwallet.CustomerID, "///", ReferenceNo, "///", User.Uid)
			if errs != nil {
				return -1, err
			}
			return rowsAffected, nil
		}
		return 1, nil

	} else if RowsAffected == 0 {
		return 0, nil
	} else {
		return -1, nil
	}

}

// ServiceGbprime implements GbprimeService.
func (g GbprimeData) ServiceGbprime(Gbprime *models.Gbprimpay) {
	repoGb := repository.NewGbprimeRepository()

	repoGb.Gbprimepay(Gbprime)
	//fmt.Println(GB)

}

func NewGbprimeDataService() GbprimeService {
	return GbprimeData{}
}
