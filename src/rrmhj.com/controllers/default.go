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

	business.LoginedUserInfo(&this.Controller)

	this.TplNames = "product/loading.tpl"
}

type PhoneController struct {
	beego.Controller
}

func (this *PhoneController) Get() {
	this.Data["IsLogin"] = business.CheckLogin(this.GetSession)
	business.LoginedUserInfo(&this.Controller)

	this.TplNames = "phone.tpl"
}

type DownController struct {
	beego.Controller
}

func (this *DownController) Get() {
	this.TplNames = "down.tpl"
}
