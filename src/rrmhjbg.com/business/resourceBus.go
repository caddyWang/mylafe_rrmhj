package business

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"rrmhjbg.com/conf"
	"rrmhjbg.com/dao"
	. "rrmhjbg.com/models/jsonmodels"
	"rrmhjbg.com/models/resource"
	"rrmhjbg.com/tools"
	"strconv"
)

var (
	roleInfoList   []resource.SrcRoleInfo
	dialogInfoList []resource.SrcDialogInfo
	sceneInfoList  []resource.SrcSceneInfo

	faceInfoList     []resource.SrcRoleFaceInfo
	actionInfoList   []resource.SrcRoleActionInfo
	clothingInfoList []resource.SrcRoleClothingInfo
)

//2013/07/23 Wangdj 新增：下载指定新角色
//2013/07/24 Wangdj 修改：将查找指定角色，表情，动作与衣服的业务代码提炼到三个公共方法
func DownNewRole(roleName, uid string) (zipbyte []byte) {

	var url = conf.ImgUrl
	var fileName []string
	var confFile []DownRes

	fileName, confFile = getRoleBySystem(roleName, 1, fileName, confFile)
	fileName, confFile, _ = getRoleFaceBySystem(roleName, 1, fileName, confFile)
	fileName, confFile, _, _ = getRoleActionClothingBySystem(roleName, 1, fileName, confFile)

	jsonRtn, err := json.Marshal(confFile)
	if err != nil {
		beego.Error("数据格式化成JSON出错！", err)
	}

	dao.SaveRoleInUser(roleName, uid)

	zipbyte = tools.GencZip(fileName, url, jsonRtn)

	return
}

func DownExistRole(roleName, uid string) (zipbyte []byte) {
	var url = conf.ImgUrl
	var fileName, faceNames, actionNames, clothingNames []string
	var confFile []DownRes

	fileName, confFile, faceNames = getRoleFaceBySystem(roleName, 0, fileName, confFile)
	fileName, confFile, actionNames, clothingNames = getRoleActionClothingBySystem(roleName, 0, fileName, confFile)

	jsonRtn, err := json.Marshal(confFile)
	if err != nil {
		beego.Error("数据格式化成JSON出错！", err)
	}

	dao.SaveRoleFaceInUser(faceNames, uid)
	dao.SaveRoleActionInUser(actionNames, uid)
	dao.SaveRoleClothingInUser(clothingNames, uid)

	zipbyte = tools.GencZip(fileName, url, jsonRtn)

	return
}

//2013/07/18 Wangdj 新增：获取当前分页下的资源信息，验证数据合法性，并转换成json
func ShowSrcInfoByPage(pageIndex, pageSize, uid, roleName string, srcType int64, showResList *ShowResList) {
	var count int
	var err, err1 error
	var downRoleInfo string

	index, size := filterPageInfo(pageIndex, pageSize)
	showResList.PageIndex = strconv.Itoa(index)
	showResList.PageSize = strconv.Itoa(size)
	showResList.ImgSuffix = conf.ImgUrl

	userDownd := resource.SrcUserDownloaded{}
	err1 = dao.GetDownloadInfoByUid(uid, &userDownd)

	switch srcType {
	case RoleType:
		roleInfoList = []resource.SrcRoleInfo{}
		count, err = dao.ShowRoleInfoByPage(index, size, &roleInfoList)

	case DialogType:
		dialogInfoList = []resource.SrcDialogInfo{}
		count, err = dao.ShowDialogInfoByPage(index, size, &dialogInfoList)

	case SceneType:
		sceneInfoList = []resource.SrcSceneInfo{}
		count, err = dao.ShowSceneInfoByPage(index, size, &sceneInfoList)

	case RoleFaceType:
		faceInfoList = []resource.SrcRoleFaceInfo{}
		count, err = dao.ShowRoleFaceInfoByPage(index, size, roleName, &faceInfoList)

	case RoleActionType:
		actionInfoList = []resource.SrcRoleActionInfo{}
		count, err = dao.ShowRoleActionInfoByPage(index, size, roleName, &actionInfoList)

	case RoleClothingType:
		clothingInfoList = []resource.SrcRoleClothingInfo{}
		count, err = dao.ShowRoleClothingInfoByPage(index, size, roleName, &clothingInfoList)
	}

	if err != nil || err1 != nil {
		showResList.OptCode = "-1"
	} else {
		showResList.OptCode = "0"
		showResList.ListCount = strconv.Itoa(count)

		switch srcType {
		case RoleType:
			for _, role := range userDownd.RoleInfo {
				downRoleInfo += " " + role.RoleName
			}
			for _, rec := range roleInfoList {
				showResList.ListArry = append(showResList.ListArry, rec.GetRes(downRoleInfo))
			}

		case DialogType:
			downRoleInfo = fmt.Sprint(userDownd.DialogInfo)
			for _, rec := range dialogInfoList {
				showResList.ListArry = append(showResList.ListArry, rec.GetRes(downRoleInfo))
			}

		case SceneType:
			downRoleInfo = fmt.Sprint(userDownd.SceneInfo)
			for _, rec := range sceneInfoList {
				showResList.ListArry = append(showResList.ListArry, rec.GetRes(downRoleInfo))
			}

		case RoleFaceType:
			for _, role := range userDownd.RoleInfo {
				if role.RoleName == roleName {
					downRoleInfo = fmt.Sprint(role.RoleFaceInfo)
				}
			}

			for _, rec := range faceInfoList {
				showResList.ListArry = append(showResList.ListArry, rec.GetRes(downRoleInfo))
			}

		case RoleActionType:
			for _, role := range userDownd.RoleInfo {
				if role.RoleName == roleName {
					downRoleInfo = fmt.Sprint(role.RoleActionInfo)
				}
			}

			for _, rec := range actionInfoList {
				showResList.ListArry = append(showResList.ListArry, rec.GetRes(downRoleInfo))
			}

		case RoleClothingType:
			for _, role := range userDownd.RoleInfo {
				if role.RoleName == roleName {
					downRoleInfo = fmt.Sprint(role.RoleClothingInfo)
				}
			}

			for _, rec := range clothingInfoList {
				showResList.ListArry = append(showResList.ListArry, rec.GetRes(downRoleInfo))
			}
		}

	}
}

//2013/07/18 Wangdj 对分页条件进行过滤
func filterPageInfo(pageIndex, pageSize string) (index, size int) {
	var err, err1 error
	index, err = strconv.Atoi(pageIndex)
	size, err1 = strconv.Atoi(pageSize)

	if err != nil || index < 1 {
		index = 1
	}

	if err1 != nil || size < 1 {
		size = conf.PageSize
	}

	return
}

func getRoleBySystem(roleName string, systemRole int, fileName []string, confFile []DownRes) ([]string, []DownRes) {
	srcRoleInfo := []resource.SrcRoleInfo{}
	dao.GetRoleBySystem(roleName, systemRole, &srcRoleInfo)
	for _, role := range srcRoleInfo {
		fileName = append(fileName, role.PicName)
		fileName = append(fileName, role.ItemPicName)

		cf := DownRes{PicName: role.PicName, SrcType: strconv.Itoa(RoleType), KeyName: roleName, ItemPicName: role.ItemPicName, Direction: strconv.Itoa(role.Direction), DefaultFace: role.DefaultFace, DefaultClothing: role.DefaultClothing}
		confFile = append(confFile, cf)
	}

	return fileName, confFile
}

func getRoleFaceBySystem(roleName string, systemRole int, fileName []string, confFile []DownRes) ([]string, []DownRes, []string) {
	srcRoleFaceInfo := []resource.SrcRoleFaceInfo{}
	faceNames := []string{}

	dao.GetRoleFaceBySystem(roleName, systemRole, &srcRoleFaceInfo)
	for _, face := range srcRoleFaceInfo {
		fileName = append(fileName, face.PicName)
		fileName = append(fileName, face.ItemPicName)
		faceNames = append(faceNames, face.FaceName)

		cf := DownRes{PicName: face.PicName, SrcType: strconv.Itoa(RoleFaceType), KeyName: face.FaceName, ItemPicName: face.ItemPicName, RoleName: roleName}
		confFile = append(confFile, cf)
	}
	return fileName, confFile, faceNames
}

func getRoleActionClothingBySystem(roleName string, systemRole int, fileName []string, confFile []DownRes) ([]string, []DownRes, []string, []string) {
	srcRoleActionInfo := []resource.SrcRoleActionInfo{}
	var actionNames, clothingNames []string

	dao.GetRoleActionClothingBySystem(roleName, systemRole, &srcRoleActionInfo)
	for _, act := range srcRoleActionInfo {
		fileName = append(fileName, act.ItemPicName)
		actionNames = append(actionNames, act.ActionName)

		for _, cl := range act.Clothing {
			fileName = append(fileName, cl.PicName)
			fileName = append(fileName, "item-"+cl.PicName)
			clothingNames = append(clothingNames, cl.ClothingName)

			cf := DownRes{PicName: cl.PicName, SrcType: strconv.Itoa(RoleClothingType), KeyName: cl.ClothingName, ItemPicName: "item-" + cl.PicName, ActionItemPicName: act.ItemPicName, RoleName: roleName, ClothingGroup: cl.ClothingName + "-" + act.ActionName, ActionGroup: cl.ClothingName + "-" + act.ActionName}
			confFile = append(confFile, cf)
		}
	}

	return fileName, confFile, actionNames, clothingNames
}
