package repository

import (
	"backEndGo/models"
	"fmt"

	"gorm.io/gorm"
)

type CourseRepository interface {
	GetCourseAll() (*[]models.Course, error)
	GetCourseByIDCoach(Cid int) (*[]models.Course, error)
	UpdateStatusCourse(CoID int, Status string) int64
	GetCouseByname(Name string) (*[]models.Course, error)
	GetCouseByCoID(CoID int) (*models.Course, error)
	UpdateCourse(course *models.Course) int64
}
type courseDB struct {
	db *gorm.DB
}

// UpdateCourse implements CourseRepository
func (c courseDB) UpdateCourse(course *models.Course) int64 {
	result := c.db.Model(models.Course{}).Where("coID = ?", course.CoID).Updates(
		models.Course{Name: course.Name, Details: course.Details, Level: course.Level, Amount: course.Amount,
			Image: course.Image, Days: course.Days, Price: course.Price, Status: course.Status})

	if result.Error != nil {
		panic(result.Error)
	}
	if result.RowsAffected > 0 {
		fmt.Println("Update completed")
	}
	return result.RowsAffected
}

// GetCoachByCoID implements CourseRepository
func (c courseDB) GetCouseByCoID(CoID int) (*models.Course, error) {
	course := models.Course{}
	result := c.db.Where("coID = ?", CoID).Find(&course)
	if result.Error != nil {
		return nil, result.Error
	}
	return &course, nil
}

// GetCouseByname implements CourseRepository
func (c courseDB) GetCouseByname(Name string) (*[]models.Course, error) {
	courses := []models.Course{}
	result := c.db.Where("name like ?", "%"+Name+"%").Where("bid IS NULL").Find(&courses)
	if result.Error != nil {
		return nil, result.Error
	}
	return &courses, nil
}

// UpdateStatusCourse implements CourseRepository
func (c courseDB) UpdateStatusCourse(CoID int, Status string) int64 {

	result := c.db.Model(models.Course{}).Where("coID = ?", CoID).Update("status", Status)
	if result.Error != nil {
		panic(result.Error)
	}
	if result.RowsAffected > 0 {
		fmt.Println("Update completed")
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
