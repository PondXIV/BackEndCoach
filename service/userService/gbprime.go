package userservice

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type GbprimeService interface {
	ServiceGbprime(Gbprime *models.Gbprimpay)
	ServiceWallet(ReferenceNo string) (int64, error)
}

type GbprimeData struct {
}

// ServiceWallet implements GbprimeService.
func (GbprimeData) ServiceWallet(ReferenceNo string) (int64, error) {
	repoWallet := repository.NewWalletRepository()
	RowsAffected, err := repoWallet.UpdateWallet(ReferenceNo)
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
