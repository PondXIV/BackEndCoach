package repository

import (
	"backEndGo/models"

	"gorm.io/gorm"
)

type DayOfCourseRepository interface {
	DayOfCourseAll() (*[]models.DayOfCouse, error)
	DayOfCourseByCoid(CoID int) (*[]models.DayOfCouse, error)
	DayOfCourseByDid(Did int) (*models.DayOfCouse, error)
	InsertDayOfCourse(CoID uint, Days int) int
}
type DayOfCourseDB struct {
	db *gorm.DB
}

// DayOfCourseByDid implements DayOfCourseRepository
func (d DayOfCourseDB) DayOfCourseByDid(Did int) (*models.DayOfCouse, error) {
	days := models.DayOfCouse{}
	result := d.db.Preload("Course").Where("did = ?", Did).Find(&days)

	if result.Error != nil {
		return nil, result.Error
	}
	return &days, nil
}

// DayOfCourseByCoid implements DayOfCourseRepository
func (d DayOfCourseDB) DayOfCourseByCoid(CoID int) (*[]models.DayOfCouse, error) {
	days := []models.DayOfCouse{}
	result := d.db.Preload("Course").Where("coID = ?", CoID).Find(&days)

	if result.Error != nil {
		return nil, result.Error
	}
	return &days, nil
}

// DayOfCourseAll implements DayOfCourseRepository
func (d DayOfCourseDB) DayOfCourseAll() (*[]models.DayOfCouse, error) {
	days := []models.DayOfCouse{}
	result := d.db.Preload("Course").Find(&days)

	if result.Error != nil {
		return nil, result.Error
	}
	return &days, nil
}

// InsertDaayOfCourse implements DayOfCourseRepository
func (d DayOfCourseDB) InsertDayOfCourse(CoID uint, Days int) int {
	num := 0
	for i := 1; i <= Days; i++ {
		dayOfCourse := models.DayOfCouse{
			CourseID: CoID,
			Sequence: i,
			//Course:   models.Course{},
		}
		result := d.db.Create(&dayOfCourse)
		num += int(result.RowsAffected)
	}

	// if result.Error != nil {
	// 	panic(result.Error)
	// }
	return num
}

func NewDayOfCourseRepository() DayOfCourseRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}

	return DayOfCourseDB{db}
}
