package buycourse

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type ShowBuyingDataService interface {
	GetBuying(uid int, coID int) (*[]models.Buying, error)
}
type BuyingData struct {
}

// SeviceGetListClipByIDCoach implements ShowListClipDataService
func (c BuyingData) GetBuying(uid int, coID int) (*[]models.Buying, error) {
	repo := repository.NewBuyingRepository()
	Buys, err := repo.GetBuyingrAll(uid, coID, 0)
	if err != nil {
		return nil, err
	}
	return Buys, nil
}

func NewBuyingDataService() ShowBuyingDataService {
	return BuyingData{}
}
