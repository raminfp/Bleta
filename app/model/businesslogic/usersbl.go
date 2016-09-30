package businesslogic

import (
	"github.com/raminfp/Bleta/app/model"
	"github.com/raminfp/Bleta/database/drivers"
)

func Insert_Users(pmodel model.Users) bool {
	db := drivers.Connection()
	defer db.Close()
	db.AutoMigrate(&model.Users{})
	db.Create(&pmodel)
	return true

}

