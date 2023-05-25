package clipincourse

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type InsertClipDataService interface {
	SeviceInsertClip(Did int, Clip *models.Clip) (int64, error)
}
type InsertClipData struct {
}

// SeviceInsertListClip implements InsertListClipDataService
func (InsertClipData) SeviceInsertClip(Did int, Clip *models.Clip) (int64, error) {
	repoClip := repository.NewClipRepository()

	getAllClip, err := repoClip.GetClip(0, 0, Did)

	if err != nil {
		panic(err)
	}
	for _, c := range *getAllClip {
		if c.ListClipID == Clip.ListClipID {
			return 0, nil
		}
	}
	RowsAffected, err := repoClip.InsertClip(Did, Clip)
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

func NewInsertClipDataService() InsertClipDataService {
	return InsertClipData{}
}
