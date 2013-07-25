package resource

import (
	"github.com/astaxie/beego"
	"rrmhjbg.com/models/jsonmodels"
	"strings"
	"time"
)

type SrcRoleFaceInfo struct {
	Id          string "_id"
	FaceName    string
	PicName     string
	ItemPicName string
	RoleName    string
	PostTime    time.Time
	SystemRole  int
	ProfileName string
	ProfileImg  string
	Iflag       int
}

func (this *SrcRoleFaceInfo) GetRes(downloadInfo string) (res jsonmodels.Res) {
	res = jsonmodels.Res{}
	res.KeyName = this.FaceName
	res.ItemPic = this.ItemPicName
	res.IsDown, res.TipNum = "0", "0"
	res.ProfileName = this.ProfileName
	res.ProfilePic = this.PicName
	res.ProfileText = ""

	//如果是系统角色或者用户已经下载过，进行标识
	beego.Debug("downloadInfo", downloadInfo)
	beego.Debug("facename", this.FaceName)
	if this.SystemRole == 1 || strings.Contains(downloadInfo, this.FaceName) {
		res.IsDown = "1"
	}

	return res
}

type SrcRoleClothingInfo struct {
	Id           string "_id"
	ClothingName string
	PicName      string
	ItemPicName  string
	RoleName     string
	PostTime     time.Time
	SystemRole   int
	ProfileName  string
	Iflag        int
}

func (this *SrcRoleClothingInfo) GetRes(downloadInfo string) (res jsonmodels.Res) {
	res = jsonmodels.Res{}
	res.KeyName = this.ClothingName
	res.ItemPic = this.ItemPicName
	res.IsDown, res.TipNum = "0", "0"
	res.ProfileName = this.ProfileName
	res.ProfilePic = this.PicName
	res.ProfileText = ""

	//如果是系统角色或者用户已经下载过，进行标识
	if this.SystemRole == 1 || strings.Contains(downloadInfo, this.ClothingName) {
		res.IsDown = "1"
	}

	return res
}

type SrcRoleActionInfo struct {
	Id         string "_id"
	ActionName string
	Clothing   []struct {
		ClothingName string
		PicName      string
	}
	ItemPicName string
	RoleName    string
	SystemRole  int
	PostTime    time.Time
	ProfileName string
	Iflag       int
}

func (this *SrcRoleActionInfo) GetRes(downloadInfo string) (res jsonmodels.Res) {
	res = jsonmodels.Res{}
	res.KeyName = this.ActionName
	res.ItemPic = this.ItemPicName
	res.IsDown, res.TipNum = "0", "0"
	res.ProfileName = this.ProfileName
	res.ProfileText = ""

	//如果是系统角色或者用户已经下载过，进行标识
	if this.SystemRole == 1 || strings.Contains(downloadInfo, this.ActionName) {
		res.IsDown = "1"
	}

	if len(this.Clothing) > 0 {
		res.ProfilePic = this.Clothing[0].PicName
	}

	return res
}
