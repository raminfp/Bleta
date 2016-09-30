package businesslogic

import (
	"github.com/raminfp/Bleta/app/model"
	"github.com/raminfp/Bleta/database/drivers"

	"github.com/jinzhu/gorm"
)

func Insert_Users(pmodel model.Users) bool {
	db := drivers.Connection()
	defer db.Close()
	db.AutoMigrate(&model.Users{})
	db.Create(&pmodel)
	return true

}

func SelectAll_Users(pmodel model.Users) ( *gorm.DB) {
	db := drivers.Connection()
	defer db.Close()
	lst := db.Find(&pmodel)
	return lst
}

func SelectOne_Users(pmodel model.Users,pID uint) (*gorm.DB)  {
	db := drivers.Connection()
	defer db.Close()
	record := db.First(&pmodel, pID)
	return record
}

func Update_Usrs() {

}