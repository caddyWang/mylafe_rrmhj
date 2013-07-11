package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"rrmhj.com/business"
	"rrmhj.com/models"
	"strings"
)

type ProCommentController struct {
	beego.Controller
}

func (this *ProCommentController) Get() {
	proid := this.GetString("pid")
	commlist, _, _ := business.GetProComment(proid)

	beego.Debug("commentlist =", commlist)

	infoJson, err1 := json.Marshal(commlist)
	if err1 != nil {
		beego.Error("数据格式化成JSON出错！", err1)
	}

	this.Ctx.WriteString(string(infoJson))
}

func (this *ProCommentController) Post() {

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

	proId, dingface := this.GetString("proId"), this.GetString("dingface")

	beego.Debug(proId)
	beego.Debug(dingface)
	this.Ctx.SetCookie(proId, "1", 0)
	business.UpdateProUporDown(proId, dingface)

	this.TplNames = "blank.tpl"
}
