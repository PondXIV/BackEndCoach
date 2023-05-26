package clipsv

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type UpdateListClipDataService interface {
	ServiceUpdateListClip(IcpID int, clip *models.ListClip) (int64, error)
}
type UpdateListClipData struct {
}

// ServiceUpdateListFood implements UpdateListFoodDataService
func (UpdateListClipData) ServiceUpdateListClip(IcpID int, clip *models.ListClip) (int64, error) {
	repoListClip := repository.NewListClipRepository()
	getAllClip, err := repoListClip.GetListClip(0, clip.CoachID, "")

	if err != nil {
		panic(err)
	}
	for _, c := range *getAllClip {

		if c.Name == clip.Name {
			if c.IcpID != uint(IcpID) {
				return 14, nil
			}

		}

	}

	RowsAffected, err := repoListClip.UpdateListClip(IcpID, clip)
	if err != nil {
		return 14, err
	}
	if RowsAffected > 0 {
		return 1, nil
	} else if RowsAffected == 0 {
		return 0, nil
	} else {
		return -1, nil
	}
}

func NewUpdateListClipDataService() UpdateListClipDataService {
	return UpdateListClipData{}
}
