package mysql

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var once sync.Once
var globalDB *gorm.DB
var err error

func ConnectDatabase() {
	dsn := "" // TODO: Read dsn from config file
	globalDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func GetDB() *gorm.DB {
	once.Do(func() {
		ConnectDatabase()
	})
	return globalDB
}
