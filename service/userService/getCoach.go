package userservice

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type GetNamecoachService interface {
	ServiceGetNameCoach(Name string) (*[]models.Coach, error)
	ServiceGetNameCoachs(Id int, Name string, Email string) (*[]models.Coach, error)
}
type CoachByNameData struct {
}

// ServiceGetNameCoachs implements GetNamecoachService
func (CoachByNameData) ServiceGetNameCoachs(Id int, Name string, Email string) (*[]models.Coach, error) {
	repo := repository.NewCoachRepository()
	coachs, err := repo.Getcoach(Id, Name, Email)
	if err != nil {
		panic(err)
	}
	return coachs, nil
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
