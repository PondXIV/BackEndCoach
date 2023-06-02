package repository

import (
	"backEndGo/models"

	"gorm.io/gorm"
)

type RequestRepository interface {
	GetRequest(RqID int, Uid int, Cid int) (*[]models.Request, error)
}
type RequestDB struct {
	db *gorm.DB
}

// GetRequest implements RequestRepository.
func (r RequestDB) GetRequest(RqID int, Uid int, Cid int) (*[]models.Request, error) {
	request := []models.Request{}
	result := r.db.Preload("")
	if RqID != 0 {
		result.Where("rqID=?", RqID)
	}
	if Uid != 0 {
		result.Where("uid=?", Uid)
	}
	if Cid != 0 {
		result.Where("cid=? AND status = 0", Cid)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	result.Find(&request)

	return &request, nil
}

func NewRequestRepository() RequestRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}

	return RequestDB{db}
}
