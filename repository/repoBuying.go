package repository

import (
	"backEndGo/models"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BuyingRepository interface {
	GetBuyingrAll(uid int, coID int, bid int, cid int, ocoID int) (*[]models.Buying, error)
	BuyCourse(CoID int, Weight int, Buying *models.Buying) (int, error)
	GetCourseByIDCusEX(Uid int) (*[]models.Buying, error)
	GetCourseByIDCus(Uid int) (*[]models.Buying, error)
	UpdateCoIDBuying(Bid int, CourseNewID int) error
}
type buyingDB struct {
	db *gorm.DB
}

// updateCoIDBuying implements BuyingRepository.
func (b buyingDB) UpdateCoIDBuying(Bid int, CourseNewID int) error {
	result := b.db.Model(models.Buying{}).Where("bid = ?", Bid).Update("coID", CourseNewID)
	if result.Error != nil {
		panic(result.Error)
	}
	return result.Error
}

// GetCourseByIDCus implements CourseRepository
func (b buyingDB) GetCourseByIDCus(Uid int) (*[]models.Buying, error) {
	dt := time.Now()
	day := dt.AddDate(0, 0, -1)
	buying := []models.Buying{}
	result := b.db.Joins("Course").Preload("Course.Coach").Where("uid=?", Uid).Where("expiration_date >= ? OR expiration_date IS NULL ", day).Find(&buying)
	if result.Error != nil {
		return nil, result.Error
	}
	return &buying, nil
}

// GetCourseByIDCusEX implements CourseRepository.
func (b buyingDB) GetCourseByIDCusEX(Uid int) (*[]models.Buying, error) {
	dt := time.Now()
	day := dt.AddDate(0, 0, -1)
	buying := []models.Buying{}
	result := b.db.Joins("Course").Preload("Course.Coach").Where("uid=?", Uid).Where("expiration_date < ?", day).Find(&buying)

	if result.Error != nil {
		return nil, result.Error
	}
	return &buying, nil
}

// BuyCourse implements BuyingRepository
func (b buyingDB) BuyCourse(OriginalID int, Weight int, Buying *models.Buying) (int, error) {

	result := b.db.Create(&models.Buying{
		Bid:         0,
		CustomerID:  Buying.CustomerID,
		BuyDateTime: Buying.BuyDateTime,
		CourseID:    uint(OriginalID),
		OriginalID:  uint(OriginalID),
		Weight:      Weight,
		//Customer:    models.Customer{},
	})
	if result.Error != nil {
		return -1, result.Error
	}
	buyingLast := models.Buying{}
	getBuyLast := b.db.Order("bid DESC").Find(&buyingLast)
	if getBuyLast.Error != nil {
		panic(getBuyLast.Error)
	}
	return int(buyingLast.Bid), nil
}

// GetBuyingrAll implements BuyingRepository
func (b buyingDB) GetBuyingrAll(uid int, coID int, bid int, cid int, ocoID int) (*[]models.Buying, error) {
	buying := []models.Buying{}
	result := b.db.Joins("Course").Preload("Customer").Preload("Course.Coach")
	if uid != 0 {
		fmt.Println(uid)
		result.Where("uid=?", uid)
	}
	if bid != 0 {
		result.Where("Buying.bid=?", bid)
	}
	if coID != 0 {
		result.Where("coID=?", coID)
	}
	if cid != 0 {
		result.Where("cid=?", cid)
	}
	if ocoID != 0 {
		fmt.Println(ocoID)
		result.Where("originalID=?", ocoID)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	result.Find(&buying)
	return &buying, nil
}

func NewBuyingRepository() BuyingRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}

	return buyingDB{db}
}
