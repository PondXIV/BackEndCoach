package coachsv

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type UpdateCoachDataService interface {
	ServiceUpdateCoach(Cid int, coach *models.Coach) (int64, error)
}
type CoachDataUpdate struct {
}

// ServiceUpdateCoach implements UpdateCoachDataService.
func (CoachDataUpdate) ServiceUpdateCoach(Cid int, coach *models.Coach) (int64, error) {
	repo := repository.NewCoachRepository()
	RowsAffected, err := repo.UpdateCoach(Cid, coach)
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

func NewUpdateCoachDataService() UpdateCoachDataService {
	return CoachDataUpdate{}
}
