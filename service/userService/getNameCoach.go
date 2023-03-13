package userservice

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type GetNamecoachService interface {
	ServiceGetNameCoach(Name string) (*[]models.Coach, error)
}
type CoachByNameData struct {
}

// ServiceGetNameCoach implements GetNamecoachService
func (CoachByNameData) ServiceGetNameCoach(Name string) (*[]models.Coach, error) {
	repo := repository.NewCoachRepository()
	coachs, err := repo.GetCoachByName(Name)
	if err != nil {
		panic(err)
	}
	return coachs, nil
}

func NewCoachByNameDataService() GetNamecoachService {
	return CoachByNameData{}
}
