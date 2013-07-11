package controllers

import (
	"github.com/astaxie/beego"
	"rrmhj.com/business"
	"rrmhj.com/conf"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Plist"], this.Data["ProCount"] = business.QueryProductsList(0, this.Ctx.Request)
	this.Data["IsLogin"] = business.CheckLogin(this.GetSession)
	this.Data["PageIndex"], this.Data["PageSize"] = 1, conf.PageSize

	business.LoginedUserInfo(&this.Controller)

	this.TplNames = "index.tpl"
}

func (this *MainController) Post() {
	pageIndex, _ := this.GetInt("pageIndex")

	beego.Debug("Loading PageIndex=", pageIndex)

	this.Data["IsLogin"] = business.CheckLogin(this.GetSession)
	this.Data["Plist"], _ = business.QueryProductsList(int(pageIndex), this.Ctx.Request)

	this.TplNames = "product/loading.tpl"
}
