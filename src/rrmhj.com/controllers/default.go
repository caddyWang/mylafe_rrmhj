package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"rrmhj.com/business"
	"rrmhj.com/conf"
	"rrmhj.com/models"
	"strconv"
	"strings"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Plist"], this.Data["ProCount"] = business.QueryProductsList(0, this.Ctx.Request)
	this.Data["IsLogin"] = business.CheckLogin(this.GetSession)
	this.Data["PageIndex"], this.Data["PageSize"] = 1, conf.PageSize

	this.TplNames = "index.tpl"
}

func (this *MainController) Post() {
	pageIndex, _ := this.GetInt("pageIndex")

	beego.Debug("Loading PageIndex=", pageIndex)

	this.Data["IsLogin"] = business.CheckLogin(this.GetSession)
	this.Data["Plist"], _ = business.QueryProductsList(int(pageIndex), this.Ctx.Request)

	this.TplNames = "product/loading.tpl"
}

type CommentController struct {
	beego.Controller
}

func (this *CommentController) Get() {
	proid := this.GetString("pid")
	commlist, _, _ := business.GetProComment(proid)

	infoJson, err1 := json.Marshal(commlist)
	if err1 != nil {
		beego.Error("数据格式化成JSON出错！", err1)
	}

	this.Ctx.WriteString(string(infoJson))
}

func (this *CommentController) Post() {

	proid, commdesc := this.GetString("proid"), this.GetString("commentdesc")

	var state struct {
		StateJson
		models.Comment
	}

	if !business.CheckLogin(this.GetSession) {
		state.StateCode, state.StateInfo = -1, "您登录过时了，再登一次呗！"
		beego.Debug("登录失效！proid=", proid)
	}

	if strings.Trim(commdesc, " ") == "" {
		state.StateCode, state.StateInfo = -1, "您多少说点啥吧！"
		beego.Debug("评论内容为空！proid=", proid)
	}

	if state.StateCode == 0 {
		comment := models.Comment{}
		comment.Proid, comment.CommentDesc = proid, commdesc
		if err := business.SaveProductComment(&comment, this.GetSession); err != nil {
			beego.Error("保存评论出错！comment=", comment, err)
			state.StateCode, state.StateInfo = -1, "保存评论出错！"
		}

		state.Comment = comment
	}

	infoJson, err1 := json.Marshal(state)
	if err1 != nil {
		beego.Error("数据格式化成JSON出错！", err1)
	}

	this.Ctx.WriteString(string(infoJson))
}

type ProOptController struct {
	beego.Controller
}

func (this *ProOptController) Get() {

	proId, optValue := this.GetString("proId"), this.GetString("optView")
	optIntVal, err := strconv.Atoi(optValue)
	if err != nil {
		beego.Error("顶踩参数不正确：optValue=", optValue, "proId=", proId, err)
		return
	}

	beego.Debug(proId)
	beego.Debug(optIntVal)
	this.Ctx.SetCookie(proId, optValue, 0)
	business.UpdateProUporDown(proId, optIntVal)

	this.TplNames = "blank.tpl"
}
