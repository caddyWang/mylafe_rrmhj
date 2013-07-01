package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"rrmhjbg.com/business"
	"rrmhjbg.com/models"
	"strconv"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Post() {
	rrmhjUid, iconURL, platformName, profileURL, userName, uid := this.GetString("rrmhjUid"), this.GetString("iconURL"), this.GetString("platformName"), this.GetString("profileURL"), this.GetString("userName"), this.GetString("usid")

	socialuser := models.SocialUserInfo{Uid: uid, UserName: userName, ProfileImg: iconURL, ProfileUrl: profileURL}
	dbUid := business.InitUserInfoBySinaWeibo(socialuser, platformName, rrmhjUid)

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

func (this *UserController) Get() {
	pageIndex, pageSize, _ := this.GetString("pageIndex"), this.GetString("pageSize"), this.GetString("accessToken")
	url := this.Ctx.Request.RemoteAddr

	ipageIndex, err := strconv.Atoi(pageIndex)
	if err != nil {
		ipageIndex = 0
	}
	ipageSize, err1 := strconv.Atoi(pageSize)
	if err1 != nil {
		ipageSize = 20
	}

	beego.Debug("RemoteAddr=", url)

	var rtn struct {
		OptCode   int
		UserCount int
		UserDatas []models.UserInfo
	}

	rtn.OptCode, rtn.UserCount, rtn.UserDatas = 0, 0, []models.UserInfo{}
	rtn.UserCount, rtn.OptCode = business.GetAllUsers(ipageIndex, ipageSize, &rtn.UserDatas)
	beego.Debug("rtn=", rtn)

	jsonRtn, err := json.Marshal(rtn)
	if err != nil {
		beego.Error("数据格式化成JSON出错！", err)
	}

	this.Ctx.WriteString(string(jsonRtn))
}
