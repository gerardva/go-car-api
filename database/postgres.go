package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gerardva/go-api/config"
	"github.com/gerardva/go-api/models"
)

var db *gorm.DB

func Init() {
	config := config.GetConfig()
	connect(config)
	migrate()
}

func GetDatabase() *gorm.DB {
	return db
}

func connect(config *config.Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
        panic(err)
    } else {
		fmt.Printf("Connected Database with host %s", config.DBHost)
    }
}

func migrate() {
	db.AutoMigrate(&models.Car{})
	fmt.Println("Migration complete")
}