package config

import (
	"todo-app/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {
    var err error
    dsn := "root:@tcp(127.0.0.1:3306)/todoapp?charset=utf8mb4&parseTime=True&loc=Local"
    DB, err = gorm.Open("mysql", dsn)
    if err != nil {
        panic("Failed to connect to database!")
    }

    DB.AutoMigrate(&models.User{}, &models.Todo{})
}