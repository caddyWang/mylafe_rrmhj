package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"rrmhjbg.com/business"
	"rrmhjbg.com/models"
	"rrmhjbg.com/tools"
	"strconv"
)

type UserController struct {
	beego.Controller
}
type FilterController struct {
	beego.Controller
}

func (this *UserController) Post() {
	rrmhjUid, iconURL, platformName, profileURL, userName, uid := this.GetString("rrmhjUid"), tools.FilterURL(this.GetString("iconURL")), this.GetString("platformName"), tools.FilterURL(this.GetString("profileURL")), this.GetString("userName"), this.GetString("usid")

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

	if uid != "-1" {
		business.InitUserDownInfo(dbUid)
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

func (this *FilterController) Get() {
	var UserDatas []models.UserInfo
	_, _ = business.GetAllUsers(1, 10000, &UserDatas)

	for _, u := range UserDatas {
		u.ProfileImg = tools.FilterURL(u.ProfileImg)

		var swb []models.SocialUserInfo
		for _, s := range u.SinaWeibo {
			s.ProfileImg = tools.FilterURL(s.ProfileImg)
			s.ProfileUrl = tools.FilterURL(s.ProfileUrl)

			swb = append(swb, s)
		}
		u.SinaWeibo = swb

		var twb []models.SocialUserInfo
		for _, s := range u.TencWeibo {
			s.ProfileImg = tools.FilterURL(s.ProfileImg)
			s.ProfileUrl = tools.FilterURL(s.ProfileUrl)

			twb = append(twb, s)
		}
		u.TencWeibo = twb

		var qq []models.SocialUserInfo
		for _, s := range u.QQZone {
			s.ProfileImg = tools.FilterURL(s.ProfileImg)
			s.ProfileUrl = tools.FilterURL(s.ProfileUrl)

			qq = append(qq, s)
		}
		u.QQZone = qq

		var rr []models.SocialUserInfo
		for _, s := range u.RenRenSNS {
			s.ProfileImg = tools.FilterURL(s.ProfileImg)
			s.ProfileUrl = tools.FilterURL(s.ProfileUrl)

			rr = append(rr, s)
		}
		u.RenRenSNS = rr

		business.UpdateUser(&u)

	}
	this.Ctx.WriteString("filter OK!")
}
