package clipincourse

import (
	"backEndGo/repository"
)

type DeleteClipDataService interface {
	SeviceDeleteClip(CpID int) (int64, error)
}

type DeleteClipData struct {
}

// SeviceGetListClipByIDCoach implements ShowListClipDataService
func (DeleteClipData) SeviceDeleteClip(CpID int) (int64, error) {
	repo := repository.NewClipRepository()
	RowsAffected, err := repo.DeleteClip(CpID)
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

func NewDeleteClipDataService() DeleteClipDataService {
	return DeleteClipData{}
}
