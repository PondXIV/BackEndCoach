package repository

import (
	"backEndGo/models"
	"fmt"

	"gorm.io/gorm"
)

type CoachRepository interface {
	GetCoachAll() (*[]models.Coach, error)
	GetCoachByID(Id int) (*[]models.Coach, error)
	GetCoachByName(Name string) (*[]models.Coach, error)
	Getcoach(Id int, Name string) (*[]models.Coach, error)
	UpdateCoach(Cid int, coach *models.Coach) (int64, error)
	UpdatePriceCoach(Cid int, Price int) (int64, error)
}

type coachDB struct {
	db *gorm.DB
}

// UpdatePriceCoach implements CoachRepository.
func (c coachDB) UpdatePriceCoach(Cid int, Price int) (int64, error) {
	result := c.db.Model(models.Coach{}).Where("cid = ?", Cid).Updates(
		models.Coach{Price: Price})

	if result.Error != nil {
		return -1, result.Error
	}
	if result.RowsAffected > 0 {
		fmt.Println("Update completed")
	}
	return result.RowsAffected, nil
}

// UpdateCoach implements CoachRepository.
func (c coachDB) UpdateCoach(Cid int, coach *models.Coach) (int64, error) {
	result := c.db.Model(models.Coach{}).Where("cid = ?", Cid).Updates(&coach)

	if result.Error != nil {
		return -1, result.Error
	}
	if result.RowsAffected > 0 {
		fmt.Println("Update completed")
	}
	return result.RowsAffected, nil
}

// Getcoach implements CoachRepository
func (c coachDB) Getcoach(Id int, Name string) (*[]models.Coach, error) {
	coachs := []models.Coach{}
	var result *gorm.DB = c.db.Find(&coachs)
	if Id != 0 {
		//result.Joins("Course").Where("cid = ?", Id).Find(&coachs)
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
