package controllers

import (
	"github.com/astaxie/beego"
	"rrmhj.com/business"
)

// 2013/07/12 Wangdj 新增：未登录用户，跳转到首页面
func noLoginToDefPage(ctx *beego.Controller) {
	if business.CheckLogin(ctx.GetSession) == false {
		ctx.Redirect("/", 302)
		return
	}
}
