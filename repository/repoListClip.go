package repository

import (
	"backEndGo/models"

	"gorm.io/gorm"
)

type ListClipRepository interface {
	GetListClipByIDCoach(Cid int) (*[]models.ListClip, error)
}
type ListClipDB struct {
	db *gorm.DB
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
