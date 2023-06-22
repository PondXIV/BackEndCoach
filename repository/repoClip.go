package repository

import (
	"backEndGo/models"
	"fmt"

	"gorm.io/gorm"
)

type ClipRepository interface {
	GetClip(CpID int, IcpID int, Did int) (*[]models.Clip, error)
	InsertClip(Did int, Clip *models.Clip) (int64, error)
	UpdateClip(CpID int, Clip *models.Clip) (int64, error)
	DeleteClip(CpID int) (int64, error)
	InsertBuyClip(Did int, IcpID int, Status string) (int64, error)
	UpdateStatusClip(CpID int, Status string) (int64, error)
}
type ClipDB struct {
	db *gorm.DB
}

// updateStatusClip implements ClipRepository.
func (c ClipDB) UpdateStatusClip(CpID int, Status string) (int64, error) {
	fmt.Println("ID1", CpID, "\t", "St1", Status)
	result := c.db.Model(models.Clip{}).Where("cpID = ?", CpID).Update("status", Status)
	if result.Error != nil {
		panic(result.Error)
	}
	if result.RowsAffected > 0 {
		fmt.Println("Update completed")
	}
	if result.RowsAffected > 0 {
		fmt.Println("Update completed")
	}
	return result.RowsAffected, nil
}

// InsertBuyClip implements ClipRepository.
func (c ClipDB) InsertBuyClip(Did int, IcpID int, Status string) (int64, error) {
	result := c.db.Create(&models.Clip{
		CpID:         0,
		ListClipID:   uint(IcpID),
		DayOfCouseID: uint(Did),
		Status:       Status,
	})
	if result.Error != nil {
		return -1, result.Error
	}
	if result.Error != nil {
		panic(result.Error)
	}
	return result.RowsAffected, nil
}

// GetClip implements ClipRepository
func (c ClipDB) GetClip(CpID int, IcpID int, Did int) (*[]models.Clip, error) {
	clips := []models.Clip{}
	result := c.db.Preload("ListClip")
	if CpID != 0 {
		result.Where("cpID=?", CpID)
	}
	if IcpID != 0 {
		result.Where("icpID=?", IcpID)
	}
	if Did != 0 {
		result.Where("did=?", Did)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	result.Find(&clips)

	return &clips, nil
}

// InsertClip implements ClipRepository
func (c ClipDB) InsertClip(Did int, Clip *models.Clip) (int64, error) {

	result := c.db.Create(&models.Clip{
		CpID:         0,
		ListClipID:   Clip.ListClipID,
		DayOfCouseID: uint(Did),
		Status:       Clip.Status,
	})
	if result.Error != nil {
		return -1, result.Error
	}
	if result.Error != nil {
		panic(result.Error)
	}

	return result.RowsAffected, nil
}

// UpdateClip implements ClipRepository
func (c ClipDB) UpdateClip(CpID int, Clip *models.Clip) (int64, error) {
	result := c.db.Model(models.Clip{}).Where("cpID = ?", CpID).Updates(
		models.Clip{
			ListClipID: Clip.ListClipID,
			Status:     Clip.Status,
		})
	return result.RowsAffected, nil
}

// DeleteClip implements ClipRepository
func (c ClipDB) DeleteClip(CpID int) (int64, error) {
	clipID := &models.Clip{CpID: uint(CpID)}
	result := c.db.Delete(clipID)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func NewClipRepository() ClipRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}

	return ClipDB{db}
}
