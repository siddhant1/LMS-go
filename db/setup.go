package db

import (
	"ilearn-backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Setup() *gorm.DB {
	dsn := "root:mysql@tcp(127.0.0.1:3306)/ilearn?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Unable to connect to db")
	}

	db.AutoMigrate(&models.Teacher{})
	return db

}
