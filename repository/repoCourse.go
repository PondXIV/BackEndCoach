package repository

import (
	"backEndGo/models"
	"fmt"
	"time"

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
	GetCourseByIDCusEX(Uid int) (*[]models.Course, error)
	InsertCourseByID(CoID int, Bid int) (int, int, int, error)
	DeleteCourse(CoID int) (int64, error)
	UpdateExpiration(CoID int, Days int) (int64, error)
	UpdateDay(CoID int, Day int) (int64, error)
	GetCourseByUser(Cid int) (*[]models.Course, error)
}
type courseDB struct {
	db *gorm.DB
}

// GetCourseByUser implements CourseRepository.
func (c courseDB) GetCourseByUser(Cid int) (*[]models.Course, error) {
	courses := []models.Course{}

	result := c.db.Preload("Coach").Preload("Buying.Customer").Where("cid = ?", Cid).Where("bid IS NOT NULL").Find(&courses)
	//
	if result.Error != nil {
		return nil, result.Error
	}

	return &courses, nil
}

// GetCourseByIDCusEX implements CourseRepository.
func (c courseDB) GetCourseByIDCusEX(Uid int) (*[]models.Course, error) {
	dt := time.Now()
	courses := []models.Course{}
	result := c.db.Joins("Buying").Where("uid=?", Uid).Where("expiration_date < ?", dt).Find(&courses)
	if result.Error != nil {
		return nil, result.Error
	}
	return &courses, nil
}

// UpdateDay implements CourseRepository.
func (c courseDB) UpdateDay(CoID int, Day int) (int64, error) {
	result := c.db.Model(models.Course{}).Where("coID = ?", CoID).Update("days", Day)
	if result.Error != nil {
		return -1, result.Error
	}

	if result.RowsAffected > 0 {
		fmt.Println("Update completed")
	}
	return result.RowsAffected, nil
}

// UpdateExpiration implements CourseRepository.
func (c courseDB) UpdateExpiration(CoID int, Days int) (int64, error) {
	dt := time.Now()

	day := dt.AddDate(0, 0, Days-1)

	//ex_date := day.Format("01-02-2006")
	// days := strconv.ParseInt(Days)
	// tm := time.Unix(days)
	fmt.Println(CoID)

	result := c.db.Model(models.Course{}).Where("coID = ?", CoID).Update("expiration_date", day)
	if result.Error != nil {
		panic(result.Error)
	}
	if result.RowsAffected > 0 {
		fmt.Println("Update completed")
	}
	if result.RowsAffected > 0 {
		fmt.Println("Update completed")
	}
	return result.RowsAffected, nil

}

// DeleteCourse implements CourseRepository.
func (c courseDB) DeleteCourse(CoID int) (int64, error) {
	coID := &models.Course{
		CoID: uint(CoID),
	}
	result := c.db.Delete(coID)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

// InsertCourseByID implements CourseRepository
func (c courseDB) InsertCourseByID(CoID int, Bid int) (int, int, int, error) {
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
		return -1, -1, -1, result.Error
	}
	course2 := models.Course{}
	getCourseLast := c.db.Order("coID DESC").Find(&course2)
	if getCourseLast.Error != nil {
		panic(getCourse.Error)
	}
	return int(course2.Price), int(course2.CoID), int(course2.Days), nil
}

// GetCourseAll implements CourseRepository
func (c courseDB) GetCourse(CoID int, Cid int, Name string) (*[]models.Course, error) {
	courses := []models.Course{}

	result := c.db.Preload("Coach")
	if CoID != 0 {
		result.Where("coID=?", CoID)
	}
	if Cid != 0 {
		result.Where("cid = ?", Cid).Where("bid IS NULL").Find(&courses)
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
	dt := time.Now()
	courses := []models.Course{}
	result := c.db.Joins("Buying").Where("uid=?", Uid).Where("expiration_date > ? OR expiration_date IS NULL ", dt).Find(&courses)
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
	mdcourse := models.Course{}
	getCourseLast := c.db.Order("coID DESC").Find(&mdcourse)
	if getCourseLast.Error != nil {
		panic(getCourseLast.Error)
	}
	return int64(mdcourse.CoID), nil
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
	result := c.db.Where("coID = ?", CoID).Where("bid IS NOT NULL").Find(&course)
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
