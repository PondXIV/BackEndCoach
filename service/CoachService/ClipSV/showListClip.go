package clipsv

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type ShowListClipDataService interface {
	SeviceGetListClipByIDCoach(Cid int) (*[]models.ListClip, error)
}
type ListClipData struct {
}

// SeviceGetListClipByIDCoach implements ShowListClipDataService
func (ListClipData) SeviceGetListClipByIDCoach(Cid int) (*[]models.ListClip, error) {
	repo := repository.NewListClipRepository()
	listClips, err := repo.GetListClipByIDCoach(Cid)
	if err != nil {
		return nil, err
	}
	return listClips, nil
}

func NewListClipDataService() ShowListClipDataService {
	return ListClipData{}
}
