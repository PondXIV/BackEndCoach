package repository

import (
	"backEndGo/models"
	"fmt"

	"gorm.io/gorm"
)

type CourseRepository interface {
	GetCourse(CoID int, Cid int, Name string) (*[]models.Course, error)
	GetCourseByIDCoach(Cid int) (*[]models.Course, error)
	UpdateStatusCourse(CoID int, Status string) int64
	GetCouseByname(Name string) (*[]models.Course, error)
	GetCouseByCoID(CoID int) (*models.Course, error)
	UpdateCourse(CoID int, course *models.Course) (int64, error)
	InsertCourse(Cid int, course *models.Course) (int64, error)
	GetCourseByIDCus(Uid int) (*[]models.Course, error)
	InsertCourseByID(CoID int, Bid int) (int, error)
}
type courseDB struct {
	db *gorm.DB
}

// InsertCourseByID implements CourseRepository
func (c courseDB) InsertCourseByID(CoID int, Bid int) (int, error) {
	course := models.Course{}
	getCourse := c.db.Where("coID = ?", CoID).Find(&course)
	if getCourse.Error != nil {
		panic(getCourse.Error)
	}

	//course.CoachID = Cid;
	result := c.db.Create(&models.Course{
		CoID:           0,
		CoachID:        course.CoachID,
		BuyingID:       uint(Bid),
		Name:           course.Name,
		Details:        course.Details,
		Level:          course.Level,
		Amount:         course.Amount,
		Image:          course.Image,
		Days:           course.Days,
		Price:          course.Price,
		Status:         course.Status,
		ExpirationDate: course.ExpirationDate,
	})
	if result.Error != nil {
		return -1, result.Error
	}

	return int(course.Price), nil
}

// GetCourseAll implements CourseRepository
func (c courseDB) GetCourse(CoID int, Cid int, Name string) (*[]models.Course, error) {
	courses := []models.Course{}
	result := c.db.Where("bid IS NULL")
	if CoID != 0 {
		result.Where("coID=?", CoID)
	}
	if Cid != 0 {
		result.Where("cid=?", Cid)
	}
	if Name != "" {
		result.Where("name like ?", "%"+Name+"%")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	result.Find(&courses)

	return &courses, nil
}

// GetCourseByIDCus implements CourseRepository
func (c courseDB) GetCourseByIDCus(Uid int) (*[]models.Course, error) {
	courses := []models.Course{}
	result := c.db.Joins("Buying").Where("uid=?", Uid).Find(&courses)
	if result.Error != nil {
		return nil, result.Error
	}
	return &courses, nil
}

// InsertCourse implements CourseRepository
func (c courseDB) InsertCourse(Cid int, course *models.Course) (int64, error) {
	course.CoID = 0

	//course.CoachID = Cid;
	result := c.db.Create(&models.Course{
		CoID:           course.CoID,
		CoachID:        Cid,
		BuyingID:       course.BuyingID,
		Name:           course.Name,
		Details:        course.Details,
		Level:          course.Level,
		Amount:         course.Amount,
		Image:          course.Image,
		Days:           course.Days,
		Price:          course.Price,
		Status:         course.Status,
		ExpirationDate: course.ExpirationDate,
		// Buying:         models.Buying{},
		// DayOfCouses:    []models.DayOfCouse{},
	})
	if result.Error != nil {
		return -1, result.Error
	}

	return result.RowsAffected, nil
}

// UpdateCourse implements CourseRepository
func (c courseDB) UpdateCourse(CoID int, course *models.Course) (int64, error) {
	result := c.db.Model(models.Course{}).Where("coID = ?", CoID).Updates(
		models.Course{Name: course.Name, Details: course.Details, Level: course.Level, Amount: course.Amount,
			Image: course.Image, Days: course.Days, Price: course.Price, Status: course.Status})

	if result.Error != nil {
		return -1, result.Error
	}
	if result.RowsAffected > 0 {
		fmt.Println("Update completed")
	}
	return result.RowsAffected, nil
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
	if result.RowsAffected > 0 {
		fmt.Println("Update completed")
	}
	return result.RowsAffected
}

// GetCourseByIDCoach implements CourseRepository
func (c courseDB) GetCourseByIDCoach(Cid int) (*[]models.Course, error) {
	courses := []models.Course{}
	result := c.db.Preload("DayOfCouses").Where("cid = ?", Cid).Where("bid IS NULL").Find(&courses)
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
