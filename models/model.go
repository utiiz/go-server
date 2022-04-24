package model

import (
	"fmt"

	"github.com/utiiz/go-server/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Setup(c config.Config) {
	var err error
	connectionString := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
	dsn := fmt.Sprintf(connectionString, c.Db.Address, c.Db.User, c.Db.Password, c.Db.Database, c.Db.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
}
