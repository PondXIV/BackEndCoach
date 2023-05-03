package clipsv

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type InsertListClipDataService interface {
	SeviceInsertListClip(Cid int, Clip *models.ListClip) (int64, error)
}
type InsertListClipData struct {
}

// SeviceInsertListClip implements InsertListClipDataService
func (InsertListClipData) SeviceInsertListClip(Cid int, Clip *models.ListClip) (int64, error) {
	repoListClip := repository.NewListClipRepository()

	getAllClip, err := repoListClip.GetListClipByIDCoach(Cid)

	if err != nil {
		panic(err)
	}
	for _, c := range *getAllClip {
		if c.Name == Clip.Name {
			return 0, nil
		}
	}
	RowsAffected, err := repoListClip.InsertListClip(Cid, Clip)
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

func NewInsertListClipDataService() InsertListClipDataService {
	return InsertListClipData{}
}
