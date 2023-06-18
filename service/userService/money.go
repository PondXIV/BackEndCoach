package userservice

import (
	"backEndGo/models"
	"backEndGo/repository"
)

var repoWallet = repository.NewWalletRepository()

type GbprimeService interface {
	ServiceGbprime(Gbprime *models.Gbprimpay)
	ServiceWallet(ReferenceNo string, ResGb *models.Gbprimpay) (int64, error)
	ServiceInsertWallet(CusID int, wallet *models.Wallet) (int64, error)
}

type GbprimeData struct {
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

	RowsAffected, err := repoWallet.UpdateWallet(ReferenceNo, ResGb)
	if err != nil {
		return -1, err
	}
	if RowsAffected > 0 {
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
