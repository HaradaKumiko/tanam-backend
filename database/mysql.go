package database

import (
	"fmt"
	"log"
	"tanam-backend/domains/model"
	"tanam-backend/helpers"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		helpers.GetConfig("DB_USERNAME"),
		helpers.GetConfig("DB_PASSWORD"),
		helpers.GetConfig("DB_HOST"),
		helpers.GetConfig("DB_PORT"),
		helpers.GetConfig("DB_NAME"),
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when creating a connection to the database: %s\n", err)
		panic(err.Error())
	}

	log.Println("connected to the database")
}

func Migrate() {
	models := []interface{}{&model.Plant{}, &model.Auth{}, &model.Biodata{}}

	err := DB.Migrator().DropTable(models...)
	if err != nil {
		log.Fatalf("failed to drop tables: %s\n", err)
	}

	err = DB.AutoMigrate(models...)

	if err != nil {
		log.Fatalf("failed to perform database migration: %s\n", err)
	}
}
