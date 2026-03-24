package main

import (
	"fmt"

	"github.com/mobin-alz/gameapp/entity"
	"github.com/mobin-alz/gameapp/repository/mysql"
)

func main() {
	mysqlRepo := mysql.New()
	response, err := mysqlRepo.RegisterUser(entity.User{
		PhoneNumber: "09015037617",
		Name:        "Mobin Alizadeh",
	})
	if err != nil {
		_ = fmt.Errorf("error : %v\n", err)
	} else {
		fmt.Println("User created :", response)
	}

	isUnique, err := mysqlRepo.IsPhoneNumberUnique(response.PhoneNumber)
	if err != nil {
		_ = fmt.Errorf("error : %v\n", err)
	} else {
		fmt.Println(isUnique)
	}
}
