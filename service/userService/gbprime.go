package userservice

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type GbprimeService interface {
	ServiceGbprime(Gbprime *models.Gbprimpay)
}

type GbprimeData struct {
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
