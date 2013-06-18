package controllers

import (
	"github.com/astaxie/beego"
	"rrmhj.com/business"
	"rrmhj.com/conf"
)

//每页条数
var pageSize = conf.PageSize

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Plist"], _ = business.QueryProductsList(1, pageSize)
	this.Data["IsLogin"] = business.CheckLogin(this.GetSession)

	this.TplNames = "index.tpl"
}
