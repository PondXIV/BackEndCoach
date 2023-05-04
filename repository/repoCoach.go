package repository

import (
	"backEndGo/models"

	"gorm.io/gorm"
)

type CoachRepository interface {
	GetCoachAll() (*[]models.Coach, error)
	GetCoachByID(Id int) (*[]models.Coach, error)
	GetCoachByName(Name string) (*[]models.Coach, error)
	Getcoach(Id int, Name string) (*[]models.Coach, error)
}

type coachDB struct {
	db *gorm.DB
}

// Getcoach implements CoachRepository
func (c coachDB) Getcoach(Id int, Name string) (*[]models.Coach, error) {
	coachs := []models.Coach{}
	var result *gorm.DB = c.db.Find(&coachs)
	if Id != 0 {
		result.Where("cid = ?", Id).Find(&coachs)
	}
	if Name != "" {
		result.Where("username  like ?", "%"+Name+"%").Find(&coachs)
	}
	result.Find(&coachs)
	return &coachs, nil
}

// GetCoachByName implements CoachRepository
func (c coachDB) GetCoachByName(Name string) (*[]models.Coach, error) {
	coachs := []models.Coach{}
	resultCoa := c.db.Where("username  like ?", "%"+Name+"%").Find(&coachs)
	if resultCoa.Error != nil {
		return nil, resultCoa.Error
	}
	return &coachs, nil
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

func NewCoachRepository() CoachRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}

	return coachDB{db}
}
