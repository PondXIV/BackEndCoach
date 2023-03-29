package repository

import (
	"backEndGo/models"

	"gorm.io/gorm"
)

type DayOfCourseRepository interface {
	InsertDayOfCourse(CoID uint, Days int) int
}
type DaayOfCourseDB struct {
	db *gorm.DB
}

// InsertDaayOfCourse implements DayOfCourseRepository
func (d DaayOfCourseDB) InsertDayOfCourse(CoID uint, Days int) int {
	num := 0
	for i := 1; i <= Days; i++ {
		dayOfCourse := models.DayOfCouse{
			CoID:     CoID,
			Sequence: i,
			Course:   models.Course{},
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

	return DaayOfCourseDB{db}
}
