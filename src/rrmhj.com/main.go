package main

import (
	"github.com/astaxie/beego"
	"rrmhj.com/business"
	"rrmhj.com/controllers"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/sinalogin", &controllers.SinaLoginController{})
	beego.Router("/tenclogin", &controllers.TencLoginController{})
	beego.Router("/pro/comment", &controllers.ProCommentController{})
	beego.Router("/pro/updown", &controllers.ProOptController{})
	beego.Router("/pro/like", &controllers.ProLikeController{})

	beego.AddFuncMap("fmtHeadImg", business.DefaultHeadImg)
	beego.AddFuncMap("loginDisplay", business.LoginDisplay)
	beego.AddFuncMap("logoutDisplay", business.LogoutDisplay)
	beego.AddFuncMap("displayLike", business.DisplayLike)
	beego.AddFuncMap("islike", business.IsLike)

	beego.Run()
}
