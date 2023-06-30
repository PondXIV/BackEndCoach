package requestservice

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type InsertRequestDataService interface {
	ServiceInsertRequest(CusID int, request *models.Request) (int64, error)
}
type InsertRequestData struct {
}

// ServiceInsertRequest implements InsertRequestDataService.
func (r InsertRequestData) ServiceInsertRequest(CusID int, request *models.Request) (int64, error) {
	rowsAffecteds, err := repository.NewRequestRepository().InsertRequest(CusID, request)
	if err != nil {
		panic(err)
	}
	return int64(rowsAffecteds), nil
}

func NewInsertRequestDataService() InsertRequestDataService {
	return InsertRequestData{}
}
