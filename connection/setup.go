package connection

import (
	"be-golang/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	// Koneksi ke database MySQL
	database, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/begolang?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrasi menggunakan model dari package models
	database.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Post{},
		&models.Menu{},
		&models.Sosmed{},
		&models.Dataweb{},
		&models.TokenBlacklist{},
	)

	// Menyimpan instance koneksi database ke variabel global
	DB = database
}
