package repository

import (
	"backEndGo/models"

	"gorm.io/gorm"
)

type ReviewRepository interface {
	GetReviewByIDCourse(CoID int) (*[]models.Review, error)
}
type reviewDB struct {
	db *gorm.DB
}

// GetReviewByIDCourse implements ReviewRepository
func (c reviewDB) GetReviewByIDCourse(CoID int) (*[]models.Review, error) {
	review := []models.Review{}
	result := c.db.Preload("Customer").Where("coID = ?", CoID).Find(&review)
	if result.Error != nil {
		return nil, result.Error
	}
	return &review, nil
}

func NewReviewRepository() ReviewRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}

	return reviewDB{db}
}
