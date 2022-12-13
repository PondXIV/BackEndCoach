package repository

import (
	"backEndGo/models"

	"gorm.io/gorm"
)

type CourseRepository interface {
	GetCourseAll() (*[]models.Course, error)
}
type courseDB struct {
	db *gorm.DB
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

func NewCourseRepository(gormdb *gorm.DB) CourseRepository {
	return courseDB{db: gormdb}
}
