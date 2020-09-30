package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LunchDB() error {
	err := connectToDB()
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&User{}, &Meme{}, &Comment{}, &Tag{})
	if err != nil {
		return err
	}
	return nil
}

func connectToDB() error {
	dsn := "user=postgres password=password dbname=exercise port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	return nil
}
