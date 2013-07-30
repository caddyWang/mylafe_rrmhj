package resource

import (
	"rrmhjbg.com/models/jsonmodels"
	"strings"
	"time"
)

type GtnJsonRes interface {
	GtnRes(downloadedInfo string) (res jsonmodels.Res)
}

type SrcRoleInfo struct {
	Id              string "_id"
	RoleName        string
	PicName         string
	ItemPicName     string
	MtPicName       string
	Direction       int
	DefaultFace     string
	DefaultClothing string
	SystemRole      int
	PostTime        time.Time
	ProfileName     string
	ProfileImg      string
	UndownImg       string
	ProfileText     string
	Sort            int
	Iflag           int
}

func (this *SrcRoleInfo) GetRes(downloadedInfo string) (res jsonmodels.Res) {
	res = jsonmodels.Res{}
	res.KeyName = this.RoleName
	res.ItemPic = this.MtPicName
	res.IsDown, res.TipNum = "0", "0"
	res.ProfileName = this.ProfileName
	res.ProfilePic = this.UndownImg
	res.ProfileText = this.ProfileText

	//如果是系统角色或者用户已经下载过，进行标识
	if this.SystemRole == 1 || strings.Contains(downloadedInfo, this.RoleName) {
		res.IsDown = "1"
		res.ProfilePic = this.ProfileImg
	}

	return res
}

type SrcDialogInfo struct {
	Id          string "_id"
	DialogName  string
	PicName     string
	ItemPicName string
	MtPicName   string
	Direction   int
	Color       string
	PostTime    time.Time
	SystemRole  int
	ProfileImg  string
	ProfileText string
	Sort        int
	Iflag       int
}

func (this *SrcDialogInfo) GetRes(downloadInfo string) (res jsonmodels.Res) {
	res = jsonmodels.Res{}
	res.KeyName = this.DialogName
	res.ItemPic = this.MtPicName
	res.IsDown, res.TipNum = "0", "0"
	res.ProfilePic = this.ProfileImg
	res.ProfileText = this.ProfileText

	//如果是系统角色或者用户已经下载过，进行标识
	if this.SystemRole == 1 || strings.Contains(downloadInfo, this.DialogName) {
		res.IsDown = "1"
	}

	return res
}

type SrcSceneInfo struct {
	Id          string "_id"
	SceneName   string
	PicName     string
	ItemPicName string
	MtPicName   string
	ProfileName string
	ProfileText string
	ProfileImg  string
	SystemRole  int
	PostTime    time.Time
	Sort        int
	Iflag       int
}

func (this *SrcSceneInfo) GetRes(downloadInfo string) (res jsonmodels.Res) {
	res = jsonmodels.Res{}
	res.KeyName = this.SceneName
	res.ItemPic = this.MtPicName
	res.IsDown, res.TipNum = "0", "0"
	res.ProfileName = this.ProfileName
	res.ProfilePic = this.ProfileImg
	res.ProfileText = this.ProfileText

	//如果是系统角色或者用户已经下载过，进行标识
	if this.SystemRole == 1 || strings.Contains(downloadInfo, this.SceneName) {
		res.IsDown = "1"
	}

	return res
}
