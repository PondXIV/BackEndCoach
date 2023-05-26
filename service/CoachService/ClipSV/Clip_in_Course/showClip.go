package clipincourse

import (
	"backEndGo/models"
	"backEndGo/repository"
)

type ShowClipDataService interface {
	GetClip(CpID int, IcpID int, Did int) (*[]models.Clip, error)
}
type ClipData struct {
}

// SeviceGetListClipByIDCoach implements ShowListClipDataService
func (c ClipData) GetClip(CpID int, IcpID int, Did int) (*[]models.Clip, error) {
	repo := repository.NewClipRepository()
	Clips, err := repo.GetClip(CpID, IcpID, Did)
	if err != nil {
		return nil, err
	}
	return Clips, nil
}

func NewClipDataService() ShowClipDataService {
	return ClipData{}
}
