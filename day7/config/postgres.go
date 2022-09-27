package config

import (
	// "os"
	"fmt"
	"gorm/entity"
	// "github.com/virhanali/user-management/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/joho/godotenv/autoload"
)

var DB *gorm.DB
var err error

func Database() {
	dsn := "host=localhost user=postgres password=hishna25 dbname=eFishery port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("database connect")
}

func Migrate() {
	DB.AutoMigrate(&entity.User{})
}