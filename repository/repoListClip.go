package repository

import (
	"backEndGo/models"

	"gorm.io/gorm"
)

type ListClipRepository interface {
	GetListClipByIDCoach(Cid int) (*[]models.ListClip, error)
	InsertListClip(Cid int, Clip *models.ListClip) (int64, error)
	UpdateListClip(IcpID int, Clip *models.ListClip) (int64, error)
}
type ListClipDB struct {
	db *gorm.DB
}

// UpdateListClip implements ListClipRepository
func (l ListClipDB) UpdateListClip(IcpID int, Clip *models.ListClip) (int64, error) {
	result := l.db.Model(models.ListClip{}).Where("icpID = ?", IcpID).Updates(
		models.ListClip{
			Name:         Clip.Name,
			AmountPerSet: Clip.AmountPerSet,
			Video:        Clip.Video,
			Details:      Clip.Details,
		})
	return result.RowsAffected, nil
}

// InsertListClip implements ListClipRepository
func (l ListClipDB) InsertListClip(Cid int, Clip *models.ListClip) (int64, error) {
	Clip.IcpID = 0
	result := l.db.Create(&models.ListClip{
		IcpID:        Clip.IcpID,
		CoachID:      Cid,
		Name:         Clip.Name,
		AmountPerSet: Clip.AmountPerSet,
		Video:        Clip.Video,
		Details:      Clip.Details,
	})
	if result.Error != nil {
		return -1, result.Error
	}
	if result.Error != nil {
		panic(result.Error)
	}

	return result.RowsAffected, nil
}

// GetListClipByIDCoach implements ListClipRepository
func (l ListClipDB) GetListClipByIDCoach(Cid int) (*[]models.ListClip, error) {
	clips := []models.ListClip{}
	result := l.db.Where("cid = ?", Cid).Find(&clips)

	if result.Error != nil {
		return nil, result.Error
	}
	return &clips, nil
}

func NewListClipRepository() ListClipRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}

	return ListClipDB{db}
}
