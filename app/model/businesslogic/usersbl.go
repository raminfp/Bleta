package businesslogic

import (
	"github.com/raminfp/Bleta/app/model"
	"github.com/raminfp/Bleta/database/drivers"
)


//
//db := drivers.Connection()
//defer db.Close()
//user := Users{Firstname: "ramin", Lastname: "faraj"}
//db.Create(&user)

func Insert_Users(pmodel model.Users) bool {
	db := drivers.Connection()
	defer db.Close()
	db.AutoMigrate(&model.Users{})
	db.Create(&pmodel)
	return true

}

//def Insert_UsersLogs(self,pModel):
//rep = Repository()
//return rep.Insert(pModel)
//
//def Update_UsersLogs(self,pModel):
//rep = Repository()
//return rep.Update(pModel)
//
//def SelectOne_UsersLogs(self,pID):
//rep = Repository()
//return rep.SelectOne(Logs,pID)
//
//def SelectAll_UsersLogs(self):
//rep = Repository()
//return rep.SelectAll(Logs)


