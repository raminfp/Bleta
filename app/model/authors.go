package model

import "github.com/jinzhu/gorm"

type Authors struct {
	gorm.Model
	ID        uint `gorm:"primary_key"`
	FirstName string
	LastName  string
	Email     string
}

