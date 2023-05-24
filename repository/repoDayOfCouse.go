package repository

import (
	"backEndGo/models"

	"gorm.io/gorm"
)

type DayOfCourseRepository interface {
	DayOfCourseAll() (*[]models.DayOfCouse, error)
	DayOfCourseByCoid(CoID int) (*[]models.DayOfCouse, error)
	DayOfCourseByDid(Did int) (*models.DayOfCouse, error)
	DayOfCourse(Did int, CoID int, Sequence int) (*[]models.DayOfCouse, error)
	InsertDayOfCourse(CourseID uint, Days int) int
}
type DayOfCourseDB struct {
	db *gorm.DB
}

// DayOfCourse implements DayOfCourseRepository
func (d DayOfCourseDB) DayOfCourse(Did int, CoID int, Sequence int) (*[]models.DayOfCouse, error) {
	days := []models.DayOfCouse{}
	result := d.db.Preload("Foods").Preload("Clips").Order("sequence")
	if Did != 0 {
		result.Where("did = ?", Did)
	}
	if CoID != 0 {
		result.Where("coID = ?", CoID)
	}
	if Sequence != 0 {
		result.Where("sequence = ?", Sequence).Joins("Foods")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	result.Find(&days)

	return &days, nil
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
	result := d.db.Find(&days)

	if result.Error != nil {
		return nil, result.Error
	}
	return &days, nil
}

// InsertDaayOfCourse implements DayOfCourseRepository
func (d DayOfCourseDB) InsertDayOfCourse(CourseID uint, Days int) int {
	num := 0
	for i := 1; i <= Days; i++ {
		dayOfCourse := models.DayOfCouse{
			CourseID: CourseID,
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
