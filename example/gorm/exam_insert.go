package gorm

import ("fmt"
	"github.com/raminfp/Bleta/app/model/businesslogic"
	"github.com/raminfp/Bleta/app/model"
)

func example_Insert_userModel() {

	mod := model.Users{Firstname: "ramin", Lastname: "faraj",Email:"ramin.blackhat@gmail.com",Password:"123456"}
	test := businesslogic.Insert_Users(mod)
	if test == true{
		fmt.Printf("Done")
	}


}

