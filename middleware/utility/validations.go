package utility

import (
	"github.com/asaskevich/govalidator"
)


// Email Validation

func EmailValidation(email string) bool  {

	if govalidator.IsEmail(email){
		return true
	}
	return false
}

// URL validiotn

func URLValidarion(url string) bool  {

	if govalidator.IsURL(url){
		return true
	}
	return false
}


