package daysv

import (
	"backEndGo/models"
	"backEndGo/repository"
	"fmt"
)

type ShowDayDataService interface {
	SeviceGetDay(Did int, CoID int, Sequence int) (*[]models.DayOfCouse, error)
	SeviceShowProgess(CoID int) float64
}
type DayData struct {
}

// SeviceShowProgess implements ShowDayDataService.
func (DayData) SeviceShowProgess(CoID int) float64 {
	all := 0
	sum := 0
	avg := 0.00
	//repobill := repository.NewBuyingRepository()
	repoday := repository.NewDayOfCourseRepository()
	repoclip := repository.NewClipRepository()
	days, err := repoday.DayOfCourse(0, CoID, 0)
	for _, valueDay := range *days {
		clip, _ := repoclip.GetClip(0, 0, int(valueDay.Did))
		for _, valueClip := range *clip {
			all++
			if valueClip.Status == "1" {
				sum++
			}

		}
	}
	fmt.Println("all ", all)
	fmt.Println("sum ", sum)

	avg = (float64(sum) / float64(all)) * 100
	fmt.Println("avg ", avg)
	if err != nil {
		panic(err)
	}
	return avg
}

// GetFoodByIDCoach implements ShowListFoodDataService
func (DayData) SeviceGetDay(Did int, CoID int, Sequence int) (*[]models.DayOfCouse, error) {
	repo := repository.NewDayOfCourseRepository()
	dyas, err := repo.DayOfCourse(Did, CoID, Sequence)
	if err != nil {
		return nil, err
	}
	return dyas, nil
}

func NewDayDataService() ShowDayDataService {
	return DayData{}
}
