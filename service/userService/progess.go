package userservice

import (
	"backEndGo/models"
	"backEndGo/repository"
	"fmt"
)

type ProgessbarService interface {
	//ServiceUpdateStatus(CpID int, Status string) (int64, error)
	ServiceProgess(CoID int) (*[]models.DayOfCouse, error)
}
type ProgessbarData struct {
}

// ServiceProgess implements ProgessbarService.
func (ProgessbarData) ServiceProgess(CoID int) (*[]models.DayOfCouse, error) {
	var repodays = repository.NewDayOfCourseRepository()
	var repoclip = repository.NewClipRepository()

	getday, err := repodays.DayOfCourseByCoid(CoID)
	if err != nil {
		panic(err)
	}
	for _, valueDay := range *getday {
		getclip, _ := repoclip.GetClip(0, 0, int(valueDay.Did))
		fmt.Println("Mygetclip", getclip)
	}
	return getday, nil
}

//	func (ClipData) ServiceUpdateStatus(CpID int, Status string) (int64, error) {
//		repo := repository.NewClipRepository()
//		rowsAffected, err := repo.UpdateStatusClip(CpID, Status)
//		fmt.Println("ID2", "%f", CpID, "\t", "St2", "%f", Status)
//		if err != nil {
//			panic(err)
//		}
//		return int64(rowsAffected), nil
//	}
func NewProgessbarDataService() ProgessbarService {
	return ProgessbarData{}
}
