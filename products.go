package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Db *gorm.DB

const connectionString = "Your connection string"

type Products struct {
	gorm.Model
	Name        string
	Description string
	Price       float32
}

func ConnectDatabase() error {
	var err error
	Db, err = gorm.Open("postgres", connectionString)
	if err != nil {
		return fmt.Errorf("error in connectDatabase(): %v", err)
	}
	Db.AutoMigrate(&Products{})
	return err
}
