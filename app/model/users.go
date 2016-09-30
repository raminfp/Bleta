package model

import "github.com/jinzhu/gorm"

type Users struct {
	gorm.Model
	ID        uint `gorm:"primary_key"`
	Firstname string
	Lastname  string
	Email     string
	Password  string
}
