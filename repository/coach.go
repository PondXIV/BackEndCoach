package repository

import (
	"backEndGo/models"

	"gorm.io/gorm"
)

type CoachRepository interface {
	GetCoachAll() (*[]models.Coach, error)
	LoginCoach(Email string, Password string) (*[]models.Coach, error)
}

type coachDB struct {
	db *gorm.DB
}

// LoginCoach implements CoachRepository
func (c coachDB) LoginCoach(Email string, Password string) (*[]models.Coach, error) {
	coachs := []models.Coach{}
	result := c.db.Where("email = ?", Email).Where("password = ?", Password).Find(&coachs)
	if result.Error != nil {
		return nil, result.Error
	}
	return &coachs, nil
}

// GetCoachAll implements CoachRepository
func (c coachDB) GetCoachAll() (*[]models.Coach, error) {
	coachs := []models.Coach{}
	result := c.db.Find(&coachs)
	if result.Error != nil {
		return nil, result.Error
	}
	return &coachs, nil
}

func NewCoachRepository(gormdb *gorm.DB) CoachRepository {
	return coachDB{db: gormdb}
}
