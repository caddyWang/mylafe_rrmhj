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
	MtPicName   string
	RoleName    string
	PostTime    time.Time
	SystemRole  int
	ProfileName string
	ProfileImg  string
	Sort        int
	Iflag       int
}

func (this *SrcRoleFaceInfo) GetRes(downloadInfo string) (res jsonmodels.Res) {
	res = jsonmodels.Res{}
	res.KeyName = this.FaceName
	res.ItemPic = this.MtPicName
	res.IsDown, res.TipNum = "0", "0"
	res.ProfileName = this.ProfileName
	res.ProfilePic = this.ProfileImg
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
	Id            string "_id"
	ClothingName  string
	PicName       string
	ItemPicName   string
	MtPicName     string
	RoleName      string
	ClothingGroup string
	PostTime      time.Time
	SystemRole    int
	ProfileName   string
	ProfileImg    string
	Sort          int
	Iflag         int
}

func (this *SrcRoleClothingInfo) GetRes(downloadInfo string) (res jsonmodels.Res) {
	res = jsonmodels.Res{}
	res.KeyName = this.ClothingName
	res.ItemPic = this.MtPicName
	res.IsDown, res.TipNum = "0", "0"
	res.ProfileName = this.ProfileName
	res.ProfilePic = this.ProfileImg
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
		ClothingName  string
		PicName       string
		ItemPicName   string
		ClothingGroup string
	}
	ItemPicName string
	MtPicName   string
	RoleName    string
	ActionGroup string
	SystemRole  int
	PostTime    time.Time
	ProfileName string
	ProfileImg  string
	Sort        int
	Iflag       int
}

func (this *SrcRoleActionInfo) GetRes(downloadInfo string) (res jsonmodels.Res) {
	res = jsonmodels.Res{}
	res.KeyName = this.ActionName
	res.ItemPic = this.MtPicName
	res.IsDown, res.TipNum = "0", "0"
	res.ProfileName = this.ProfileName
	res.ProfileText = ""
	res.ProfilePic = this.ProfileImg

	//如果是系统角色或者用户已经下载过，进行标识
	if this.SystemRole == 1 || strings.Contains(downloadInfo, this.ActionName) {
		res.IsDown = "1"
	}

	return res
}
