package repository

import (
	"backEndGo/models"

	"gorm.io/gorm"
)

type RequestRepository interface {
	GetRequest(RqID int, Uid int, Cid int) (*[]models.Request, error)
	InsertRequest(CusID int, request *models.Request) (int64, error)
}
type RequestDB struct {
	db *gorm.DB
}

// InsertRequest implements RequestRepository.
func (r RequestDB) InsertRequest(CusID int, request *models.Request) (int64, error) {
	result := r.db.Create(&models.Request{
		RpID:       0,
		CoachID:    request.CoachID,
		CustomerID: CusID,
		ClipID:     request.ClipID,
		Status:     "0",
		Details:    request.Details,
	})
	if result.Error != nil {
		return -1, result.Error
	}
	return result.RowsAffected, nil
}

// GetRequest implements RequestRepository.
func (r RequestDB) GetRequest(RqID int, Uid int, Cid int) (*[]models.Request, error) {
	request := []models.Request{}
	result := r.db.Preload("Customer").Preload("Clip.ListClip")

	if RqID != 0 {
		result.Where("rqID=?", RqID)
	}
	if Uid != 0 {
		result.Where("uid=?", Uid)
	}
	if Cid != 0 {
		result.Where("Request.cid=? AND status = 0", Cid)
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
