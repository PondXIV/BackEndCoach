package requestservice

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type ShowRequestDataService interface {
	ServiceGetRequest(RqID int, Uid int, Cid int) (*[]models.Request, error)
}
type RequestData struct {
}

// ServiceGetRequest implements ShowRequestDataService.
func (RequestData) ServiceGetRequest(RqID int, Uid int, Cid int) (*[]models.Request, error) {
	repo := repository.NewRequestRepository()
	requests, err := repo.GetRequest(RqID, Uid, Cid)
	if err != nil {
		return nil, err
	}
	return requests, nil
}

func NewShowRequestDataService() ShowRequestDataService {
	return RequestData{}
}
