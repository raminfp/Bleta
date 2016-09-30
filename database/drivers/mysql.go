package drivers


import (
	_ "database/sql"
        _ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/raminfp/Bleta/database/config"
	"fmt"
)

func Connection() (*gorm.DB)  {

	con := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",config.Username, config.Password,config.IpAddress,config.Port,config.DBName)
	db, err := gorm.Open("mysql", con)
	if err != nil {
		fmt.Println("Connection Error into database")
	}
	return db
}

