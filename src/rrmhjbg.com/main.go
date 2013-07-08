package main

import (
	"github.com/astaxie/beego"
	"rrmhjbg.com/controllers"
)

func main() {
	beego.Router("/user/binding_social_account", &controllers.UserController{})
	beego.Router("/user/get_all_account", &controllers.UserController{})
	beego.Router("/user/filter_user", &controllers.FilterController{})

	beego.Run()
}
