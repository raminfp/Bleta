package businesslogic

import (
	"github.com/raminfp/Bleta/app/model"
	"github.com/raminfp/Bleta/database/drivers"
)

func Insert_Author(pmodel model.Authors) bool {
	db := drivers.Connection()
	defer db.Close()
	db.AutoMigrate(&model.Authors{})
	db.Create(&pmodel)
	return true
}
