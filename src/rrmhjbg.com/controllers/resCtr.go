package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"rrmhjbg.com/business"
	"rrmhjbg.com/models/jsonmodels"
)

type ShowListController struct {
	beego.Controller
}

func (this *ShowListController) Get() {
	rrmhjUid, srcType, pageIndex, pageSize := this.GetString("rrmhjUid"), this.GetString("srcType"), this.GetString("pageIndex"), this.GetString("pageSize")

	result := jsonmodels.ShowResList{OptCode: "-1", SrcType: srcType, ListArry: []jsonmodels.Res{}}
	business.ShowSrcInfoByPage(pageIndex, pageSize, srcType, rrmhjUid, &result)

	jsonRtn, err := json.Marshal(result)
	if err != nil {
		beego.Error("数据格式化成JSON出错！", err)
	}

	this.Ctx.WriteString(string(jsonRtn))
}
