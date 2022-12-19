package repository

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func NewDatabaseConnection() (*gorm.DB, error) {
	if db == nil {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
		dsn := viper.GetString("mysql.dsn")
		dialector := mysql.Open(dsn)
		db, err = gorm.Open(dialector)
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}
