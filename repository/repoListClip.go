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
func (ListClipDB) GetListClipByIDCoach(Cid int) (*[]models.ListClip, error) {
	panic("unimplemented")
}

func NewListClipRepository() ListClipRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}

	return ListClipDB{db}
}
