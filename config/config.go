package config

import (
	"fmt"
	"yakz/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "myuser"
	password = "mypassword"
	dbname   = "mydatabase"
)

var data *gorm.DB

func ConfigDataBase() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(" failed to connect database ")
	}

	fmt.Println("✅ Connected to Database!")

	data = db

	db.AutoMigrate(&models.User{}, &models.Book{}, &models.CartBook{})

}

func GetDB() *gorm.DB {
	if data == nil {
		panic("❌ Database is not initialized! Call ConfigDataBase() first.")
	}
	return data
}
