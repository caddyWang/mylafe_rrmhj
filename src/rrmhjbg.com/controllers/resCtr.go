package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"rrmhjbg.com/business"
	. "rrmhjbg.com/models/jsonmodels"
	"strconv"
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

////////////////////////////////////////////////////////////////////////////////////////////////////

type ShowRoleInfoController struct {
	beego.Controller
}

func (this *ShowRoleInfoController) Post() {
	uid, roleName := this.GetString("rrmhjUid"), this.GetString("roleName")
	showRoleInfo := ShowRoleInfo{}
	business.GetRoleInfo(roleName, uid, &showRoleInfo)

	jsonRtn, err := json.Marshal(showRoleInfo)
	if err != nil {
		beego.Error("数据格式化成JSON出错！", err)
	}

	this.Ctx.WriteString(string(jsonRtn))
}

func (this *ShowRoleInfoController) Get() {
	this.Post()
}

////////////////////////////////////////////////////////////////////////////////////////////////////

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
	var zipName string
	switch srcType {
	case RoleType:
		if isFlag == 1 {
			zipByte, zipName = business.DownNewRole(keyName, uid)
		} else if isFlag == 2 {
			zipByte, zipName = business.DownExistRole(keyName, uid)
		}

	case RoleFaceType:
		zipByte, zipName = business.DownSingleFace(keyName, uid)

	case RoleActionType:
		zipByte, zipName = business.DownSingleAction(keyName, uid)

	case RoleClothingType:
		zipByte, zipName = business.DownSingleClothing(keyName, uid)

	case DialogType:
		zipByte, zipName = business.DownSingleDialog(keyName, uid)

	case SceneType:
		zipByte, zipName = business.DownSingleScene(keyName, uid)
	}

	if len(zipByte) > 0 {
		this.Ctx.SetHeader("Content-Length", strconv.Itoa(len(zipByte)), true)
		this.Ctx.SetHeader("Content-Type", "application/octet-stream", true)
		this.Ctx.SetHeader("Content-disposition", "attachment; filename="+zipName+".zip", true)
		this.Ctx.ResponseWriter.Write(zipByte)
	}

}

func (this *DownResController) Get() {
	if beego.RunMode == "dev" {
		this.Post()
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type RecordUserDownInfoController struct {
	beego.Controller
}

func (this *RecordUserDownInfoController) Post() {
	fileName := this.GetString("fileName")
	tipNum := business.RecordUserDownInfo(fileName)
	tip := strconv.Itoa(tipNum)

	this.Ctx.WriteString("{\"tipNum\":\"" + tip + "\"}")

}
func (this *RecordUserDownInfoController) Get() {
	if beego.RunMode == "dev" {
		this.Post()
	}
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
