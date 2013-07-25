package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"rrmhjbg.com/business"
	. "rrmhjbg.com/models/jsonmodels"
	"strconv"
	"time"
)

type ShowListController struct {
	beego.Controller
}

func (this *ShowListController) Get() {
	bindCtxData(&this.Controller)
}

func (this *ShowListController) Post() {
	bindCtxData(&this.Controller)
}

type ShowRoleListController struct {
	beego.Controller
}

func (this *ShowRoleListController) Get() {
	bindCtxData(&this.Controller)
}
func (this *ShowRoleListController) Post() {
	bindCtxData(&this.Controller)
}

type DownResController struct {
	beego.Controller
}

func (this *DownResController) Post() {

	keyName := this.GetString("keyName")
	uid := this.GetString("rrmhjUid")
	srcType, err := this.GetInt("srcType")
	if err != nil {
		beego.Error("srcType error :", err)
		return
	}
	isFlag, err1 := this.GetInt("isFlag")
	if err1 != nil {
		isFlag = 0
	}

	var zipByte []byte
	switch srcType {
	case RoleType:
		if isFlag == 1 {
			zipByte = business.DownNewRole(keyName, uid)
		} else if isFlag == 2 {
			zipByte = business.DownExistRole(keyName, uid)
		}

	case RoleFaceType:
		zipByte = business.DownSingleFace(keyName, uid)

	case RoleActionType:
		zipByte = business.DownSingleAction(keyName, uid)

	case RoleClothingType:
		zipByte = business.DownSingleClothing(keyName, uid)

	case DialogType:
		zipByte = business.DownSingleDialog(keyName, uid)

	case SceneType:
		zipByte = business.DownSingleScene(keyName, uid)
	}

	if len(zipByte) > 0 {
		fileName := strconv.FormatInt(time.Now().Unix(), 10)
		this.Ctx.SetHeader("Content-Length", strconv.Itoa(len(zipByte)), true)
		this.Ctx.SetHeader("Content-Type", "application/octet-stream", true)
		this.Ctx.SetHeader("Content-disposition", "attachment; filename="+fileName+".zip", true)
		this.Ctx.ResponseWriter.Write(zipByte)
	}

}

func (this *DownResController) Get() {
	this.Post()
}

func bindCtxData(this *beego.Controller) {
	rrmhjUid, pageIndex, pageSize := this.GetString("rrmhjUid"), this.GetString("pageIndex"), this.GetString("pageSize")
	roleName := this.GetString("roleName")

	srcType, err := this.GetInt("srcType")

	result := ShowResList{OptCode: "-1", SrcType: strconv.FormatInt(srcType, 10), ListArry: []Res{}}

	if rrmhjUid != "" || err != nil {

		if srcType < 10 || (srcType > 10 && roleName != "") {
			business.ShowSrcInfoByPage(pageIndex, pageSize, rrmhjUid, roleName, srcType, &result)
		}
	}

	jsonRtn, err := json.Marshal(result)
	if err != nil {
		beego.Error("数据格式化成JSON出错！", err)
	}

	this.Ctx.WriteString(string(jsonRtn))
}
