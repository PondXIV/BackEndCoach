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

func NewListClipRepository() ListFoodRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}

	return LisFoodDB{db}
}
