package config

import (
	"github.com/rachmankamil/kampus-merdeka-b/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	connectionString := "root@tcp(127.0.0.1:3306)/testgormDB?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func MigrateDB() {
	DB.AutoMigrate(&models.Article{})
}
