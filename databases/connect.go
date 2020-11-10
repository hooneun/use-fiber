package databases

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/hooneun/use-fiber/config"
	"github.com/hooneun/use-fiber/models"
)

// ConnectDB connect db
func ConnectDB() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), config.Config("DB_PORT"), config.Config("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	DB = db

	fmt.Println("Connect Database")
	DB.AutoMigrate(&models.User{})
	fmt.Println("Database Migrate")
}
