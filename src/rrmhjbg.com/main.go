package main

import (
	"github.com/astaxie/beego"
	"rrmhjbg.com/controllers"
)

func main() {
	beego.Router("/user/binding_social_account", &controllers.UserController{})

	beego.Run()
}
