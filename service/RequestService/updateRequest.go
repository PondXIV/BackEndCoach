package requestservice

import (
	"backEndGo/repository"
)

type UpdateRequestDataService interface {
	ServiceUpdateRequest(RqID int) (int64, error)
}
type UpdateRequestData struct {
}

// ServiceUpdateListFood implements UpdateListFoodDataService
func (UpdateRequestData) ServiceUpdateRequest(RqID int) (int64, error) {
	repo := repository.NewRequestRepository()

	RowsAffected, err := repo.UpdateRequest(RqID)
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

func NewUpdateRequestDataService() UpdateRequestDataService {
	return UpdateRequestData{}
}
