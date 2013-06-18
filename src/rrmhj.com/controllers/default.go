package controllers

import (
	"github.com/astaxie/beego"
	"rrmhj.com/business"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Plist"] = business.QueryProductsList(0)
	this.Data["IsLogin"] = business.CheckLogin(this.GetSession)

	this.TplNames = "index.tpl"
}
