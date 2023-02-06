package repository

import (
	"backEndGo/models"

	"gorm.io/gorm"
)

type CourseRepository interface {
	GetCourseAll() (*[]models.Course, error)
	GetCourseByIDCoach(Id int) (*[]models.Course, error)
}
type courseDB struct {
	db *gorm.DB
}

// GetCourseByIDCoach implements CourseRepository
func (c courseDB) GetCourseByIDCoach(Id int) (*[]models.Course, error) {
	courses := []models.Course{}
	result := c.db.Where("cid = ?", Id).Find(&courses)
	if result.Error != nil {
		return nil, result.Error
	}
	return &courses, nil
}

// GetCourseAll implements CourseRepository
func (c courseDB) GetCourseAll() (*[]models.Course, error) {
	courses := []models.Course{}
	result := c.db.Preload("Coach").Find(&courses)
	if result.Error != nil {
		return nil, result.Error
	}
	return &courses, nil
}

func NewCourseRepository() CourseRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}

	return courseDB{db}
}
