package buycourse

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type ShowBuyingDataService interface {
	GetBuying(uid int, coID int, cid int, ocoID int) (*[]models.Buying, error)
	SeviceGetCourseByUser(Cid int) (*[]models.Buying, error)
}
type BuyingData struct {
}

// SeviceGetCourseByUser implements ShowCourseDataService.
func (b BuyingData) SeviceGetCourseByUser(Cid int) (*[]models.Buying, error) {
	repo := repository.NewBuyingRepository()
	course, err := repo.GetBuyingrAll(0, 0, 0, Cid, 0)
	keys := make(map[int]bool)
	buying := []models.Buying{}
	for _, entry := range *course {
		if _, value := keys[int(entry.CustomerID)]; !value {
			keys[int(entry.CustomerID)] = true
			buying = append(buying, entry)
		}
	}
	if err != nil {
		panic(err)
	}
	return &buying, nil
}

// SeviceGetListClipByIDCoach implements ShowListClipDataService
func (c BuyingData) GetBuying(uid int, coID int, cid int, ocoID int) (*[]models.Buying, error) {
	repo := repository.NewBuyingRepository()
	Buys, err := repo.GetBuyingrAll(uid, coID, 0, cid, ocoID)
	if err != nil {
		return nil, err
	}
	return Buys, nil
}

func NewBuyingDataService() ShowBuyingDataService {
	return BuyingData{}
}
