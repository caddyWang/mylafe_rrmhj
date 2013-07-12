package controllers

import (
	"github.com/astaxie/beego"
	"rrmhj.com/business"
	"rrmhj.com/conf"
)

// 2013/07/10 Wangdj 新增：用户收藏作品Ajax请求页面
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

type MyProController struct {
	beego.Controller
}

//2013/07/11 Wangdj 新增：我发布的作品页面
func (this *MyProController) Get() {
	noLoginToDefPage(&this.Controller)

	business.LoginedUserInfo(&this.Controller)
	this.Data["MyPro"] = true
	this.Data["PageIndex"], this.Data["PageSize"] = 1, conf.PageSize

	proList, icount := business.GetProductsByUid(&this.Controller, 0)
	if icount == 0 {
		this.Data["ListNull"] = true
	}
	this.Data["Plist"], this.Data["ProCount"] = proList, icount
	_, this.Data["LikeCount"] = business.GetLikeProByUid(&this.Controller, 0)

	this.TplNames = "member/mycenter.tpl"
}

type ProDelController struct {
	beego.Controller
}

// 2013/07/12 Wangdj 新增：用户删除当前作品Ajax请求页面
func (this *ProDelController) Get() {
	proId := this.GetString("proId")

	if business.CheckLogin(this.GetSession) == false {
		this.Ctx.WriteString("-2")
		return
	}

	uid := this.GetSession("uid")
	if proId != "" && uid != "" {
		err := business.DelProductByUid(uid.(string), proId)
		if err != nil {
			this.Ctx.WriteString("-1")
		} else {
			this.Ctx.WriteString("0")
		}

		return
	}

	this.Ctx.WriteString("-1")
}

//2013/07/12 Wangdj 新增：拖动到浏览器底部时，自动加载下一页作品的ajax后台函数
func (this *MyProController) Post() {

	noLoginToDefPage(&this.Controller)
	pageIndex, _ := this.GetInt("pageIndex")

	beego.Debug("Loading PageIndex=", pageIndex)

	this.Data["IsLogin"] = true
	this.Data["Plist"], _ = business.GetProductsByUid(&this.Controller, int(pageIndex))

	business.LoginedUserInfo(&this.Controller)

	this.TplNames = "product/loading.tpl"
}

//2013/07/12 Wangdj 新增：退出登录
type ExitController struct {
	beego.Controller
}

func (this *ExitController) Get() {
	business.Logout(&this.Controller)

	this.Redirect("/", 302)
}
