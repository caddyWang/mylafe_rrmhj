package business

import (
	"fmt"
	"github.com/astaxie/beego"
	"rrmhjbg.com/business/cache"
	"rrmhjbg.com/conf"
	"rrmhjbg.com/dao"
	. "rrmhjbg.com/models/jsonmodels"
	"rrmhjbg.com/models/resource"
	"rrmhjbg.com/tools"
	"strconv"
	"strings"
	"time"
)

var (
	roleInfoList   []resource.SrcRoleInfo
	dialogInfoList []resource.SrcDialogInfo
	sceneInfoList  []resource.SrcSceneInfo

	faceInfoList     []resource.SrcRoleFaceInfo
	actionInfoList   []resource.SrcRoleActionInfo
	clothingInfoList []resource.SrcRoleClothingInfo

	url      = conf.ImgUrl
	srcCache = cache.Cache{}
)

type userDownInfo struct {
	Uid      string
	RoleName string
	SrcInfo  interface{}
}

//2013/08/12 Wangdj 新增：检测服务器上记录的最新版本信息
func DetectNewVersion(currentVer int) (ver VersionInfo) {
	ver = VersionInfo{}
	ver.HasNewVer = "0"

	verNewest, err := beego.AppConfig.Int("ver_newest")
	if err != nil {
		return
	}

	if currentVer >= verNewest {
		return
	}

	ver.HasNewVer = "1"
	ver.ImgPrefix = beego.AppConfig.String("ver_img_prefix")
	ver.ImgTip = beego.AppConfig.String("ver_img_tip")
	ver.VerInt = beego.AppConfig.String("ver_newest")

	verText := beego.AppConfig.String("ver_text")
	ver.VerText = strings.Split(verText, "|")
	ver.VerAndroidDown = beego.AppConfig.String("ver_android_down")
	ver.VerIosDown = beego.AppConfig.String("ver_ios_down")

	return
}

//2013/07/26 Wangdj 新增：清空当前用户的下载记录
func InitUserDownInfo(uid string) {
	dao.InitUserDownInfo(uid)
}

func GetRoleInfo(roleName, uid string, showRole *ShowRoleInfo) {
	srcRole := resource.SrcRoleInfo{}
	tipNum := dao.GetRoleTipNum(roleName, uid)
	dao.GetRoleInfo(roleName, &srcRole)

	if srcRole.RoleName == "" {
		showRole.OptCode = "-1"
		return
	}

	showRole.OptCode = "0"
	showRole.KeyName = srcRole.RoleName
	showRole.ProfileName = srcRole.ProfileName
	showRole.ProfilePic = srcRole.ProfileImg
	showRole.ProfileText = srcRole.ProfileText
	showRole.ImgSuffix = conf.ImgUrl
	showRole.TipNum = strconv.Itoa(tipNum)
}

//2013/07/26 Wangdj 新增：记录当前用户下载记录
//2013/08/02 Wangdj 修改：记录下载角色相关资源（表情、动作、服装或全部）时，返回当前角色的tipNum值
func RecordUserDownInfo(fileName string) (newSrcTipNum int) {
	keys := strings.Split(fileName, "_")
	if len(keys) < 2 {
		beego.Error("[rrmhjbg.com/business/RecordUserDownInfo(fileName=", fileName, ")]记录下载的文件标识不正确，如‘234343434_1’")
		return
	}

	srcType, err := strconv.Atoi(keys[1])
	if err != nil {
		beego.Error("[rrmhjbg.com/business/RecordUserDownInfo(fileName=", fileName, ")]文件标识下划线后必须为数字，如‘234343434_1’")
		return
	}

	switch srcType {
	case RoleType:
		cacheVal := srcCache.Get(fileName).(userDownInfo)
		dao.SaveRoleInUser(cacheVal.SrcInfo.(string), cacheVal.Uid)
		srcCache.Del(fileName)
		newSrcTipNum = dao.GetRoleTipNum(cacheVal.RoleName, cacheVal.Uid)

	case RoleFaceType:
		faceKey := keys[0] + "_" + strconv.Itoa(RoleFaceType)
		cacheVal := srcCache.Get(faceKey).(userDownInfo)
		dao.SaveRoleFaceInUser(cacheVal.SrcInfo.([]string), cacheVal.Uid)
		srcCache.Del(faceKey)
		newSrcTipNum = dao.GetRoleTipNum(cacheVal.RoleName, cacheVal.Uid)
		if len(keys) == 4 {
			actionKey, clothingKey := keys[0]+"_"+strconv.Itoa(RoleActionType), keys[0]+"_"+strconv.Itoa(RoleClothingType)
			actionCache := srcCache.Get(actionKey).(userDownInfo)
			clothingCache := srcCache.Get(clothingKey).(userDownInfo)
			dao.SaveRoleActionInUser(actionCache.SrcInfo.([]string), actionCache.Uid)
			dao.SaveRoleClothingInUser(clothingCache.SrcInfo.([]string), clothingCache.Uid)
			srcCache.Del(actionKey)
			srcCache.Del(clothingKey)
			newSrcTipNum = 0
		}

	case RoleActionType:
		cacheVal := srcCache.Get(fileName).(userDownInfo)
		dao.SaveRoleActionInUser(cacheVal.SrcInfo.([]string), cacheVal.Uid)
		srcCache.Del(fileName)
		newSrcTipNum = dao.GetRoleTipNum(cacheVal.RoleName, cacheVal.Uid)

	case RoleClothingType:
		cacheVal := srcCache.Get(fileName).(userDownInfo)
		dao.SaveRoleClothingInUser(cacheVal.SrcInfo.([]string), cacheVal.Uid)
		srcCache.Del(fileName)
		newSrcTipNum = dao.GetRoleTipNum(cacheVal.RoleName, cacheVal.Uid)

	case DialogType:
		cacheVal := srcCache.Get(fileName).(userDownInfo)
		dao.SaveDialogInUser(cacheVal.SrcInfo.([]string), cacheVal.Uid)
		srcCache.Del(fileName)

	case SceneType:
		cacheVal := srcCache.Get(fileName).(userDownInfo)
		dao.SaveSceneInUser(cacheVal.SrcInfo.([]string), cacheVal.Uid)
		srcCache.Del(fileName)
	}

	return
}

//2013/07/23 Wangdj 新增：下载指定新角色
//2013/07/24 Wangdj 修改：将查找指定角色，表情，动作与衣服的业务代码提炼到三个公共方法
func DownNewRole(roleName, uid string) (zipbyte []byte, zipName string) {
	var fileName []string
	var confFile []DownRes

	fileName, confFile = getRoleBySystem(roleName, 0, fileName, confFile)
	fileName, confFile, _ = getRoleFaceBySystem(roleName, 1, fileName, confFile)
	fileName, confFile, _, _ = getRoleActionClothingBySystem(roleName, 1, fileName, confFile)

	zipName = strconv.FormatInt(time.Now().Unix(), 10) + "_" + strconv.Itoa(RoleType)
	jsonRtn := tools.TransformJSON(NewDownRes{FileName: zipName, ImgStruct: confFile})
	zipbyte = tools.GencZip(fileName, url, jsonRtn)

	srcCache.Put(zipName, userDownInfo{uid, roleName, roleName})

	return
}

//2013/07/24 Wangdj 新增：下载指定角色非系统（默认）资源
func DownExistRole(roleName, uid string) (zipbyte []byte, zipName string) {
	var fileName, faceNames, actionNames, clothingNames []string
	var confFile []DownRes

	fileName, confFile, faceNames = getRoleFaceBySystem(roleName, 0, fileName, confFile)
	fileName, confFile, actionNames, clothingNames = getRoleActionClothingBySystem(roleName, -1, fileName, confFile)

	zipName = strconv.FormatInt(time.Now().Unix(), 10)
	srcCache.Put(zipName+"_"+strconv.Itoa(RoleFaceType), userDownInfo{uid, roleName, faceNames})
	srcCache.Put(zipName+"_"+strconv.Itoa(RoleActionType), userDownInfo{uid, roleName, actionNames})
	srcCache.Put(zipName+"_"+strconv.Itoa(RoleClothingType), userDownInfo{uid, roleName, clothingNames})

	zipName = zipName + "_" + strconv.Itoa(RoleFaceType) + "_" + strconv.Itoa(RoleActionType) + "_" + strconv.Itoa(RoleClothingType)
	jsonRtn := tools.TransformJSON(NewDownRes{FileName: zipName, ImgStruct: confFile})
	zipbyte = tools.GencZip(fileName, url, jsonRtn)

	return
}

//2013/07/25 Wangdj 新增：下载指定单个表情
func DownSingleFace(faceName, uid string) (zipbyte []byte, zipName string) {

	face := resource.SrcRoleFaceInfo{}
	isExist := dao.GetOneRoleFaceByKey(faceName, &face)

	if isExist {

		cf := DownRes{PicName: face.PicName, SrcType: strconv.Itoa(RoleFaceType), KeyName: face.FaceName, ItemPicName: face.ItemPicName, RoleName: face.RoleName}

		zipName = strconv.FormatInt(time.Now().Unix(), 10) + "_" + strconv.Itoa(RoleFaceType)
		srcCache.Put(zipName, userDownInfo{uid, face.RoleName, []string{face.FaceName}})
		jsonRtn := tools.TransformJSON(NewDownRes{FileName: zipName, ImgStruct: []DownRes{cf}})
		zipbyte = tools.GencZip([]string{face.PicName, face.ItemPicName}, url, jsonRtn)
	}

	return
}

//2013/07/25 Wangdj 新增：下载指定单个动作
//2013/07/31 Wangdj 增加：只下载当前动作下，此用户已下载的服装图片。以修正下载此动作，所有服装图片都下载的bug
func DownSingleAction(actionName, uid string) (zipbyte []byte, zipName string) {

	action := resource.SrcRoleActionInfo{}
	isExist := dao.GetOneRoleActionByKey(actionName, &action)

	if isExist {

		fileName := []string{action.ItemPicName}
		var confFile []DownRes

		downClothings := ""
		userDownd := resource.SrcUserDownloaded{}
		err := dao.GetDownloadInfoByUid(uid, &userDownd)
		if err == nil {
			downClothings = fmt.Sprint(userDownd.RoleClothingInfo)
		}

		//查询当前角色下的所有服装，看哪些是默认安装的 (systemrole=1)
		systemRoleMap := beego.NewBeeMap()
		clothinInfos := []resource.SrcRoleClothingInfo{}
		_, err = dao.ShowRoleClothingInfoByPage(1, 9999, action.RoleName, &clothinInfos)
		if err != nil {
			return
		}

		for _, cl := range clothinInfos {
			systemRoleMap.Set(cl.ClothingGroup, cl.SystemRole)

			if strings.Contains(downClothings, cl.ClothingName) {
				systemRoleMap.Set(cl.ClothingGroup, 1)
			}
		}

		for _, clothing := range action.Clothing {
			//2013/07/31 Wangdj
			if systemRoleMap.Get(clothing.ClothingGroup) == 1 {
				fileName = append(fileName, clothing.PicName)
				fileName = append(fileName, clothing.ItemPicName)

				cf := DownRes{PicName: clothing.PicName, SrcType: strconv.Itoa(RoleClothingType), KeyName: clothing.ClothingName, ItemPicName: clothing.ItemPicName, ActionItemPicName: action.ItemPicName, RoleName: action.RoleName, ClothingGroup: clothing.ClothingGroup, ActionGroup: action.ActionGroup}
				confFile = append(confFile, cf)
			}
		}

		zipName = strconv.FormatInt(time.Now().Unix(), 10) + "_" + strconv.Itoa(RoleActionType)
		jsonRtn := tools.TransformJSON(NewDownRes{FileName: zipName, ImgStruct: confFile})
		srcCache.Put(zipName, userDownInfo{uid, action.RoleName, []string{action.ActionName}})
		zipbyte = tools.GencZip(fileName, url, jsonRtn)
	}

	return
}

//2013/07/25 Wangdj 新增：下载指定单件衣服
func DownSingleClothing(clothingName, uid string) (zipbyte []byte, zipName string) {

	action := []resource.SrcRoleActionInfo{}
	clothing := resource.SrcRoleClothingInfo{}
	isExist := dao.GetOneRoleClothingByKey(clothingName, &clothing, &action)

	if isExist {

		downAction := ""
		userDownd := resource.SrcUserDownloaded{}
		err := dao.GetDownloadInfoByUid(uid, &userDownd)
		if err == nil {
			downAction = fmt.Sprint(userDownd.RoleActionInfo)
		}

		fileName := []string{clothing.ItemPicName}
		var confFile []DownRes

		for _, at := range action {
			//fileName = append(fileName, at.ItemPicName)

			if at.SystemRole == 1 || strings.Contains(downAction, at.ActionName) {
				for _, cl := range at.Clothing {
					if cl.ClothingGroup == clothing.ClothingGroup {
						fileName = append(fileName, cl.PicName)
						cf := DownRes{PicName: cl.PicName, SrcType: strconv.Itoa(RoleClothingType), KeyName: cl.ClothingName, ItemPicName: clothing.ItemPicName, ActionItemPicName: at.ItemPicName, RoleName: at.RoleName, ClothingGroup: cl.ClothingGroup, ActionGroup: at.ActionGroup}
						confFile = append(confFile, cf)
					}
				}
			}
		}

		zipName = strconv.FormatInt(time.Now().Unix(), 10) + "_" + strconv.Itoa(RoleClothingType)
		jsonRtn := tools.TransformJSON(NewDownRes{FileName: zipName, ImgStruct: confFile})
		srcCache.Put(zipName, userDownInfo{uid, clothing.RoleName, []string{clothing.ClothingName}})

		zipbyte = tools.GencZip(fileName, url, jsonRtn)
	}

	return
}

//2013/07/25 Wangdj 新增：下载指定单个表情
func DownSingleDialog(dialogName, uid string) (zipbyte []byte, zipName string) {

	dialog := resource.SrcDialogInfo{}
	isExist := dao.GetOneDialogByKey(dialogName, &dialog)

	if isExist {

		cf := DownRes{PicName: dialog.PicName, SrcType: strconv.Itoa(DialogType), KeyName: dialog.DialogName, ItemPicName: dialog.ItemPicName, Direction: strconv.Itoa(dialog.Direction), Color: dialog.Color}

		zipName = strconv.FormatInt(time.Now().Unix(), 10) + "_" + strconv.Itoa(DialogType)
		jsonRtn := tools.TransformJSON(NewDownRes{FileName: zipName, ImgStruct: []DownRes{cf}})
		srcCache.Put(zipName, userDownInfo{uid, "", []string{dialog.DialogName}})

		zipbyte = tools.GencZip([]string{dialog.PicName, dialog.ItemPicName}, url, jsonRtn)
	}

	return
}

//2013/07/25 Wangdj 新增：下载指定单个表情
func DownSingleScene(sceneName, uid string) (zipbyte []byte, zipName string) {

	scene := resource.SrcSceneInfo{}
	isExist := dao.GetOneSceneByKey(sceneName, &scene)

	if isExist {

		cf := DownRes{PicName: scene.PicName, SrcType: strconv.Itoa(SceneType), KeyName: scene.SceneName, ItemPicName: scene.ItemPicName}

		zipName = strconv.FormatInt(time.Now().Unix(), 10) + "_" + strconv.Itoa(SceneType)
		jsonRtn := tools.TransformJSON(NewDownRes{FileName: zipName, ImgStruct: []DownRes{cf}})
		srcCache.Put(zipName, userDownInfo{uid, "", []string{scene.SceneName}})

		zipbyte = tools.GencZip([]string{scene.PicName, scene.ItemPicName}, url, jsonRtn)
	}

	return
}

//2013/07/18 Wangdj 新增：获取当前分页下的资源信息，验证数据合法性，并转换成json
//2013/07/30 Wangdj 修改：查找当前用户下指定角色有多少新增资源
func ShowSrcInfoByPage(pageIndex, pageSize, uid, roleName string, srcType int64, showResList *ShowResList) {
	var count int
	var err, err1 error
	var downRoleInfo string

	index, size := filterPageInfo(pageIndex, pageSize)
	/*cacheKey := cache.SrcShowList + strconv.FormatInt(srcType, 10) + "_" + strconv.Itoa(index) + "_" + strconv.Itoa(size) + "_RN-" + roleName + "_UI-" + uid
	if srcCache.IsExist(cacheKey) {
		*showResList = srcCache.Get(cacheKey).(ShowResList)
		beego.Error(cacheKey)
		return
	}*/

	showResList.PageIndex = strconv.Itoa(index)
	showResList.PageSize = strconv.Itoa(size)
	showResList.ImgSuffix = conf.ImgUrl

	userDownd := resource.SrcUserDownloaded{}
	err1 = dao.GetDownloadInfoByUid(uid, &userDownd)

	switch srcType {
	case RoleType:
		roleInfoList = []resource.SrcRoleInfo{}
		count, err = dao.ShowRoleInfoByPage(index, size, &roleInfoList)
		//count, err, roleInfoList = ShowRoleInfoByPageWithCache(index, size)

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
			downRoleInfo = fmt.Sprint(userDownd.RoleInfo)
			for _, rec := range roleInfoList {
				jsonRec := rec.GetRes(downRoleInfo)
				jsonRec.TipNum = strconv.Itoa(dao.GetRoleTipNum(rec.RoleName, uid)) // 查找当前用户下当前角色有多少新增资源
				showResList.ListArry = append(showResList.ListArry, jsonRec)
			}

			//add cache
			//srcCache.Put(cacheKey, *showResList)

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
			downRoleInfo = fmt.Sprint(userDownd.RoleFaceInfo)
			for _, rec := range faceInfoList {
				showResList.ListArry = append(showResList.ListArry, rec.GetRes(downRoleInfo))
			}

		case RoleActionType:
			downRoleInfo = fmt.Sprint(userDownd.RoleActionInfo)
			for _, rec := range actionInfoList {
				showResList.ListArry = append(showResList.ListArry, rec.GetRes(downRoleInfo))
			}

		case RoleClothingType:
			downRoleInfo = fmt.Sprint(userDownd.RoleClothingInfo)
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
	existFileName := beego.NewBeeMap()

	dao.GetRoleActionClothingBySystem(roleName, systemRole, &srcRoleActionInfo)
	for _, act := range srcRoleActionInfo {
		fileName = append(fileName, act.ItemPicName)
		actionNames = append(actionNames, act.ActionName)

		for _, cl := range act.Clothing {
			if !existFileName.Check(cl.PicName) {
				fileName = append(fileName, cl.PicName)
				existFileName.Set(cl.PicName, 1)
			}
			if !existFileName.Check(cl.ItemPicName) {
				fileName = append(fileName, cl.ItemPicName)
				existFileName.Set(cl.ItemPicName, 1)
			}
			clothingNames = append(clothingNames, cl.ClothingName)

			cf := DownRes{PicName: cl.PicName, SrcType: strconv.Itoa(RoleClothingType), KeyName: cl.ClothingName, ItemPicName: cl.ItemPicName, ActionItemPicName: act.ItemPicName, RoleName: roleName, ClothingGroup: cl.ClothingGroup, ActionGroup: act.ActionGroup}
			confFile = append(confFile, cf)
		}
	}

	return fileName, confFile, actionNames, clothingNames
}

///////////////////////////////////////////////////////////////////////////////////
// Cache Opt
///////////////////////////////////////////////////////////////////////////////////

func ShowRoleInfoByPageWithCache(pageIndex, pageSize int) (count int, err error, srcRoleInfo []resource.SrcRoleInfo) {
	cacheKey := cache.SrcShowList + strconv.Itoa(RoleType) + "_" + strconv.Itoa(pageIndex) + "_" + strconv.Itoa(pageSize)
	if srcCache.IsExist(cacheKey) {
		beego.Error(cacheKey)
		srcRoleInfo = srcCache.Get(cacheKey).([]resource.SrcRoleInfo)
		count = len(srcRoleInfo)
	} else {
		srcRoleInfo = []resource.SrcRoleInfo{}
		count, err = dao.ShowRoleInfoByPage(pageIndex, pageSize, &srcRoleInfo)
		srcCache.Put(cacheKey, srcRoleInfo)
	}

	return
}

func init() {
	srcCache.Init()
}
