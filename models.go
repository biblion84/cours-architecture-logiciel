package main

import (
	"github.com/jinzhu/gorm"
)

type Products struct {
	gorm.Model
	Name        string
	Description string
	Price       float32
}
