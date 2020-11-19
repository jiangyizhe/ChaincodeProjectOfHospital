package controller

import (
	"github.com/ChainblockProject/service"
	"github.com/ChainblockProject/service2"
	"github.com/ChainblockProject/service3"
)

type Application struct {
	Setup *service.ServiceSetup
	Setup2 *service2.ServiceSetup
	Setup3 *service3.ServiceSetup
}

type User struct {
	LoginName	string
	Password	string
	IsAdmin	string
}


var users []User

func init() {

	admin := User{LoginName:"admin", Password:"123456", IsAdmin:"T"}
	alice := User{LoginName:"jiangyizhe", Password:"123456", IsAdmin:"T"}
	bob := User{LoginName:"liyue", Password:"123456", IsAdmin:"F"}
	jack := User{LoginName:"guozhongmin", Password:"123456", IsAdmin:"F"}

	users = append(users, admin)
	users = append(users, alice)
	users = append(users, bob)
	users = append(users, jack)

}