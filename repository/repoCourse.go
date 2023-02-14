package repository

import (
	"backEndGo/models"

	"gorm.io/gorm"
)

type CourseRepository interface {
	GetCourseAll() (*[]models.Course, error)
	GetCourseByIDCoach(Cid int) (*[]models.Course, error)
	UpdateStatusCourse(Id int,Status int) int64
}
type courseDB struct {
	db *gorm.DB
}

// UpdateStatusCourse implements CourseRepository
func (c courseDB) UpdateStatusCourse(Id int, Status int) int64 {
	courses := []models.Course{}
	result := c.db.Update("Course Set Status = ?",Status).Where("cid = ?", Id).Find(&courses)
	if result.Error != nil {
		panic(result.Error)
	}
	return result.RowsAffected
}
// GetCourseByIDCoach implements CourseRepository
func (c courseDB) GetCourseByIDCoach(Cid int) (*[]models.Course, error) {
	courses := []models.Course{}
	result := c.db.Where("cid = ?", Cid).Where("bid IS NULL").Find(&courses)
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
