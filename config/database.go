package config

import (
	"d_gita_be/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "dgita:dgita123@tcp(dgita-db.cbuoaypgqh0v.us-east-1.rds.amazonaws.com:3306)/dgita?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	initMigrate()
}

func initMigrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Receipt{})
	DB.AutoMigrate(&models.ImageReceipt{})
}
