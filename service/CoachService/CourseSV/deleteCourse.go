package coursesv

type DeleteCourseDataService interface {
}

type DeleteCourseData struct {
}

func NewDeleteCourseDataService() DeleteCourseDataService {
	return DeleteCourseData{}
}
