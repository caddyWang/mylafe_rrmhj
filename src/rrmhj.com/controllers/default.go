package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"rrmhj.com/business"
	"rrmhj.com/models"
	"strings"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Plist"] = business.QueryProductsList(0)
	this.Data["IsLogin"] = business.CheckLogin(this.GetSession)

	this.TplNames = "index.tpl"
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
