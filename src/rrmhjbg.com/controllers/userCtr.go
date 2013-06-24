package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"rrmhjbg.com/business"
	"rrmhjbg.com/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Post() {
	iconURL, platformName, profileURL, userName, uid := this.GetString("iconURL"), this.GetString("platformName"), this.GetString("profileURL"), this.GetString("userName"), this.GetString("usid")

	socialuser := models.SocialUserInfo{Uid: uid, UserName: userName, ProfileImg: iconURL, ProfileUrl: profileURL}
	dbUid := business.InitUserInfoBySinaWeibo(socialuser, platformName)

	var rtn struct {
		OptCode  string
		RrmhjUid string
	}
	rtn.OptCode, rtn.RrmhjUid = "0", dbUid

	if dbUid == "-1" {
		rtn.OptCode = "-1"
	}

	jsonRtn, err := json.Marshal(rtn)
	if err != nil {
		beego.Error("数据格式化成JSON出错！", err)
	}

	this.Ctx.WriteString(string(jsonRtn))
}
