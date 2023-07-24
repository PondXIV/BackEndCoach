package repository

import (
	"backEndGo/models"

	"gorm.io/gorm"
)

type ReviewRepository interface {
	GetReviewByIDCourse(CoID int) (*[]models.Review, error)
	InsertReview(CoID int, review *models.Review) (int64, error)
}
type reviewDB struct {
	db *gorm.DB
}

// InsertReview implements ReviewRepository.
func (r reviewDB) InsertReview(CoID int, review *models.Review) (int64, error) {
	result := r.db.Create(&models.Review{
		Rid:        0,
		CustomerID: review.CustomerID,
		CourseID:   CoID,
		Details:    review.Details,
		Score:      review.Score,
		Weight:     review.Weight,
		Customer:   models.Customer{},
	})
	if result.Error != nil {
		return -1, result.Error
	}
	return result.RowsAffected, nil
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
