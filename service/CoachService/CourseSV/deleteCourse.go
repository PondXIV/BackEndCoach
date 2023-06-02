package coursesv

import "backEndGo/repository"

type DeleteCourseDataService interface {
	SeviceDeleteCourse(CoID int) (int64, error)
}

type DeleteCourseData struct {
}

// SeviceDeleteCourse implements DeleteCourseDataService.
func (DeleteCourseData) SeviceDeleteCourse(CoID int) (int64, error) {
	repo := repository.NewCourseRepository()
	RowsAffected, err := repo.DeleteCourse(CoID)
	if err != nil {
		return -1, err
	}
	if RowsAffected > 0 {
		return 1, nil
	} else if RowsAffected == 0 {
		return 0, nil
	} else {
		return -1, nil
	}
}

func NewDeleteCourseDataService() DeleteCourseDataService {
	return DeleteCourseData{}
}
