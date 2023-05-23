package clipsv

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type ShowListClipDataService interface {
	SeviceGetListClip(IcpID int, Cid int, Name string) (*[]models.ListClip, error)
}
type ListClipData struct {
}

// SeviceGetListClipByIDCoach implements ShowListClipDataService
func (ListClipData) SeviceGetListClip(IcpID int, Cid int, Name string) (*[]models.ListClip, error) {
	repo := repository.NewListClipRepository()
	listClips, err := repo.GetListClip(IcpID, Cid, Name)
	if err != nil {
		return nil, err
	}
	return listClips, nil
}

func NewListClipDataService() ShowListClipDataService {
	return ListClipData{}
}
