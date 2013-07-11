package controllers

import (
	"github.com/astaxie/beego"
	"rrmhj.com/business"
)

// 2013/07/10 Wangdj 新增：用户收藏作品Ajax求页面
type ProLikeController struct {
	beego.Controller
}

func (this *ProLikeController) Get() {
	proId := this.GetString("proId")
	uid := this.GetSession("uid")

	if proId != "" && uid != "" {
		err := business.SaveUserLikeProduct(proId, uid.(string))
		if err != nil {
			this.Ctx.WriteString("-1")
		} else {
			this.Ctx.WriteString("0")
		}

		return
	}

	this.Ctx.WriteString("-1")
}

//2013/07/11 Wangdj 新增：我发布的作品页面
type MyProController struct {
	beego.Controller
}

func (this *MyProController) Get() {
	if business.CheckLogin(this.GetSession) == false {
		this.Redirect("/", 302)
		return
	}

	business.LoginedUserInfo(&this.Controller)
	this.Data["MyPro"] = true

	proList, icount := business.GetProductsByUid(&this.Controller, 0)
	if icount == 0 {
		this.Data["ListNull"] = true
	}
	this.Data["Plist"], this.Data["ListCount"] = proList, icount

	this.TplNames = "member/mycenter.tpl"
}
