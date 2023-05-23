package clipsv

import (
	"backEndGo/repository"
)

type DeleteListClipDataService interface {
	SeviceDeleteListClip(IcpID int) (int64, error)
}

type DeleteListClipData struct {
}

// SeviceGetListClipByIDCoach implements ShowListClipDataService
func (DeleteListClipData) SeviceDeleteListClip(IcpID int) (int64, error) {
	repo := repository.NewListClipRepository()
	RowsAffected, err := repo.DeleteListClip(IcpID)
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

func NewDeleteListClipDataService() DeleteListClipDataService {
	return DeleteListClipData{}
}
