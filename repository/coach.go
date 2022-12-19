package repository

import (
	"backEndGo/models"

	"gorm.io/gorm"
)

type CoachRepository interface {
	GetCoachAll() (*[]models.Coach, error)
	GetCoachByID(Id int) (*[]models.Coach, error)
}

type coachDB struct {
	db *gorm.DB
}

// GetCoachByID implements CoachRepository
func (c coachDB) GetCoachByID(Id int) (*[]models.Coach, error) {
	coachs := []models.Coach{}
	resultCoa := c.db.Where("cid = ?", Id).Find(&coachs)
	if resultCoa.Error != nil {
		return nil, resultCoa.Error
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
