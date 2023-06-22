package userservice

import (
	"backEndGo/repository"
	"fmt"
)

type ClipService interface {
	ServiceUpdateStatus(CpID int, Status string) (int64, error)
}
type ClipData struct {
}

func (ClipData) ServiceUpdateStatus(CpID int, Status string) (int64, error) {
	repo := repository.NewClipRepository()
	rowsAffected, err := repo.UpdateStatusClip(CpID, Status)
	fmt.Println("ID2", "%f", CpID, "\t", "St2", "%f", Status)
	if err != nil {
		panic(err)
	}
	return int64(rowsAffected), nil
}
func NewClipUpdateStatusDataService() ClipService {
	return ClipData{}
}
