package main

import (
	"github.com/astaxie/beego"
	"rrmhj.com/business"
	"rrmhj.com/controllers"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/sinalogin", &controllers.SinaLoginController{})
	beego.Router("/pro/comment", &controllers.CommentController{})

	beego.AddFuncMap("fmtHeadImg", business.DefaultHeadImg)
	beego.AddFuncMap("loginDisplay", business.LoginDisplay)
	beego.AddFuncMap("logoutDisplay", business.LogoutDisplay)

	beego.Run()
}
