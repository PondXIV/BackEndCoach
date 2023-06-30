package repository

import (
	"backEndGo/models"
	"fmt"

	"gorm.io/gorm"
)

type DayOfCourseRepository interface {
	DayOfCourseAll() (*[]models.DayOfCouse, error)
	DayOfCourseByCoid(CoID int) (*[]models.DayOfCouse, error)
	DayOfCourseByDid(Did int) (*models.DayOfCouse, error)
	DayOfCourse(Did int, CoID int, Sequence int) (*[]models.DayOfCouse, error)
	InsertDayOfCourse(CourseID uint, Days int) int
	BuyInsertDayOfCourse(CourseID uint, Days int) (*[]models.DayOfCouse, int)
	UpdateDay(Did int, Day *models.DayOfCouse) (int64, error)
	DeleteDay(Did int) (int64, error)
}
type DayOfCourseDB struct {
	db *gorm.DB
}

// DeleteDay implements DayOfCourseRepository.
func (d DayOfCourseDB) DeleteDay(Did int) (int64, error) {
	dayID := &models.DayOfCouse{
		Did: uint(Did),
	}
	result := d.db.Delete(dayID)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

// UpdateDay implements DayOfCourseRepository.
func (d DayOfCourseDB) UpdateDay(Did int, Day *models.DayOfCouse) (int64, error) {
	result := d.db.Model(models.DayOfCouse{}).Where("did = ?", Did).Updates(
		models.DayOfCouse{
			Sequence: Day.Sequence,
			Foods:    []models.Food{},
			Clips:    []models.Clip{},
		})

	if result.Error != nil {
		return -1, result.Error
	}
	if result.RowsAffected > 0 {
		fmt.Println("Update completed")
	}
	return result.RowsAffected, nil
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
func (d DayOfCourseDB) BuyInsertDayOfCourse(CourseID uint, Days int) (*[]models.DayOfCouse, int) {
	num := 0
	//dayslist := []uint{}

	//var daysID uint = 0
	for i := 1; i <= Days; i++ {
		dayOfCourse := models.DayOfCouse{
			CourseID: CourseID,
			Sequence: i,
		}
		result := d.db.Create(&dayOfCourse)
		num += int(result.RowsAffected)
	}
	day := []models.DayOfCouse{}
	//fmt.Printf("DIDDay = %d \t", dayslist, "\n")
	result := d.db.Where("coID = ?", CourseID).Find(&day)
	fmt.Println(&day)
	if result.Error != nil {
		panic("error")
	}

	return &day, num
}

func NewDayOfCourseRepository() DayOfCourseRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}

	return DayOfCourseDB{db}
}
