package clipincourse

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type UpdateClipDataService interface {
	ServiceUpdateClip(CpID int, Clip *models.Clip) (int64, error)
}
type UpdateClipData struct {
}

// ServiceUpdateListFood implements UpdateListFoodDataService
func (UpdateClipData) ServiceUpdateClip(CpID int, Clip *models.Clip) (int64, error) {
	repoClip := repository.NewClipRepository()
	getAllClip, err := repoClip.GetClip(0, int(Clip.ListClipID), 0)

	if err != nil {
		panic(err)
	}
	for _, c := range *getAllClip {

		if c.ListClipID == Clip.ListClipID {
			if c.CpID != uint(CpID) {
				return 3, nil
			}

		}

	}

	RowsAffected, err := repoClip.UpdateClip(CpID, Clip)
	if err != nil {
		return 3, err
	}
	if RowsAffected > 0 {
		return 1, nil
	} else if RowsAffected == 0 {
		return 0, nil
	} else {
		return -1, nil
	}
}

func NewUpdateClipDataService() UpdateClipDataService {
	return UpdateClipData{}
}
