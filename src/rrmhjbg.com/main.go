package main

import (
	"github.com/astaxie/beego"
	"rrmhjbg.com/controllers"
)

func main() {
	beego.Router("/user/binding_social_account", &controllers.UserController{})
	beego.Router("/user/get_all_account", &controllers.UserController{})
	beego.Router("/user/filter_user", &controllers.FilterController{})

	beego.Router("/src/show_list", &controllers.ShowListController{})
	beego.Router("/src/show_role_list", &controllers.ShowRoleListController{})
	beego.Router("/src/down_resource", &controllers.DownResController{})
	beego.Router("/src/init_shop", &controllers.InitUserDownInfoController{})
	beego.Router("/src/record_downinfo", &controllers.RecordUserDownInfoController{})

	beego.Router("/temp/init_data", &controllers.InitDataController{})

	beego.SetLevel(beego.LevelError)
	beego.Run()
}
