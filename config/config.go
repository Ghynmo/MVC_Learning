package config

import (
	"github.com/ghynmo/MVC_Learning/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const JWT_SECRET string = "testmvc"

func InitDB() {
	connectionString := "root@tcp(127.0.0.1:3306)/testgormDB?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func MigrateDB() {
	DB.AutoMigrate(&models.Book{})
}
