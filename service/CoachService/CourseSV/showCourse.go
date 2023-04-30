package coursesv

import (
	"backEndGo/models"
	"backEndGo/repository"
	//"time"
)

type ShowCourseDataService interface {
	ServiceGetCourseByIDCoach(Id int) (*[]models.Course, error)
	SeviceGetCourseByName(Name string) (*[]models.Course, error)
	SeviceGetCourseByCoID(CoID int) (*models.Course, error)
	SeviceGetCourseAll() (*[]models.Course, error)
}
type CourseData struct {
}

// SeviceGetCourseAll implements ShowCourseDataService
func (CourseData) SeviceGetCourseAll() (*[]models.Course, error) {
	repo := repository.NewCourseRepository()
	course, err := repo.GetCourseAll()

	if err != nil {
		panic(err)
	}
	return course, nil
}

// SeviceGetCourseByCoID implements ShowCourseDataService
func (CourseData) SeviceGetCourseByCoID(CoID int) (*models.Course, error) {
	repo := repository.NewCourseRepository()
	course, err := repo.GetCouseByCoID(CoID)
	//days, err := repository.NewDayOfCourseRepository().DayOfCourseByCoid(CoID)
	// for _, c := range *days {
	// 	// if c.CoID == models.Course[course.CoID] {
	// 	// 	// return 0, nil
	// 	// }
	// 	days, err := repository.NewDayOfCourseRepository().DayOfCourseByCoid(int(c.CoID))
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// }
	if err != nil {
		panic(err)
	}
	return course, nil
}

// SeviceGetCourseByName implements ShowCourseDataService
func (CourseData) SeviceGetCourseByName(Name string) (*[]models.Course, error) {
	repo := repository.NewCourseRepository()
	course, err := repo.GetCouseByname(Name)
	if err != nil {
		panic(err)
	}
	return course, nil
}

// ServiceGetCourseByIDCoach implements ShowCourseDataService
func (CourseData) ServiceGetCourseByIDCoach(Id int) (*[]models.Course, error) {

	// course, err := repo.GetCourseByIDCoach(Id)
	repo := repository.NewCourseRepository()
	course, err := repo.GetCourseByIDCoach(Id)

	if err != nil {
		panic(err)
	}
	return course, nil
}

func NewCourseDataService() ShowCourseDataService {
	return CourseData{}
}
