package service

import (
	// "backEndGo/models"
	// "backEndGo/repository"
	// "fmt"

	"gorm.io/gorm"
)

type GetByIDService interface {
	// GetCoaById(Id int) (*[]models.Coach, error)
}
type GetByIDData struct {
	db *gorm.DB
}

// // GetCoachByID implements GetByIDService
// func (g GetByIDData) GetCoaById(Id int) (*[]models.Coach, error) {
// 	repo := repository.NewUserRepository(g.db)
// 	coach, err := repo.GetCoachByID(Id)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return coach, nil
// }

func NewGetByIDService(gormdb *gorm.DB) GetByIDService {
	return GetByIDData{db: gormdb}
}
