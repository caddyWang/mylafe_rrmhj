package dao

/************************************************************************************
//
// Desc		:	资源库数据Dao
//
************************************************************************************/

import (
	"github.com/astaxie/beego"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"rrmhjbg.com/models/resource"
	"time"
)

const (
	roleInfo     = "srcRoleInfo"
	dialogInfo   = "srcDialogInfo"
	sceneInfo    = "srcSceneInfo"
	faceInfo     = "srcRoleFaceInfo"
	actionInfo   = "srcRoleActionInfo"
	clothingInfo = "srcRoleClothingInfo"
	userDown     = "srcUserDownloaded"

	NormalCode = 0
	DelCode    = -1
)

//2013/07/17 Wangdj 新增：根据分页获取角色列表
func ShowRoleInfoByPage(pageIndex, pageSize int, srcRoleInfo *[]resource.SrcRoleInfo) (count int, err error) {
	return showResourceInfoByPage(pageIndex, pageSize, roleInfo, "", srcRoleInfo)
}

//2013/07/18 Wangdj 新增：根据分页获取对话框列表
func ShowDialogInfoByPage(pageIndex, pageSize int, srcDialogInfo *[]resource.SrcDialogInfo) (count int, err error) {
	return showResourceInfoByPage(pageIndex, pageSize, dialogInfo, "", srcDialogInfo)
}

//2013/07/18 Wangdj 新增：根据分页获取场景列表
func ShowSceneInfoByPage(pageIndex, pageSize int, srcSceneInfo *[]resource.SrcSceneInfo) (count int, err error) {
	return showResourceInfoByPage(pageIndex, pageSize, sceneInfo, "", srcSceneInfo)
}

//2013/07/19 Wangdj 新增：根据分页获取根据分页获取相应角色表情列表
func ShowRoleFaceInfoByPage(pageIndex, pageSize int, roleName string, srcSceneInfo *[]resource.SrcRoleFaceInfo) (count int, err error) {
	return showResourceInfoByPage(pageIndex, pageSize, faceInfo, roleName, srcSceneInfo)
}

//2013/07/19 Wangdj 新增：根据分页获取根据分页获取相应角色动作列表
func ShowRoleActionInfoByPage(pageIndex, pageSize int, roleName string, srcSceneInfo *[]resource.SrcRoleActionInfo) (count int, err error) {
	return showResourceInfoByPage(pageIndex, pageSize, actionInfo, roleName, srcSceneInfo)
}

//2013/07/19 Wangdj 新增：根据分页获取根据分页获取相应角色服装列表
func ShowRoleClothingInfoByPage(pageIndex, pageSize int, roleName string, srcSceneInfo *[]resource.SrcRoleClothingInfo) (count int, err error) {
	return showResourceInfoByPage(pageIndex, pageSize, clothingInfo, roleName, srcSceneInfo)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//2013/07/30 Wangdj 新增：获取当前角色下有多少新增资源
func GetRoleTipNum(roleName, uid string) (tipNum int) {
	var faceNum, actionNum, clothingNum int

	sUDI := resource.SrcUserDownloaded{}
	err := GetDownloadInfoByUid(uid, &sUDI)
	if err != nil {
		return 0
	}

	faceNum, err = FindCount(bson.M{"rolename": roleName, "systemrole": 0, "facename": bson.M{"$not": bson.M{"$in": sUDI.RoleFaceInfo}}}, faceInfo)
	if err != nil {
		beego.Error("[rrmhjbg.com/dao/resourceDao.GetRoleTipNum(uid=", uid, "roleName=", roleName, ")] {rolename: ", roleName, ", systemrole: 0, facename: {$not: {$in: ", sUDI.RoleFaceInfo, "}}}", err)
	}
	actionNum, err = FindCount(bson.M{"rolename": roleName, "systemrole": 0, "actionname": bson.M{"$not": bson.M{"$in": sUDI.RoleActionInfo}}}, actionInfo)
	if err != nil {
		beego.Error("[rrmhjbg.com/dao/resourceDao.GetRoleTipNum(uid=", uid, "roleName=", roleName, ")] {rolename: ", roleName, ", systemrole: 0, actionname: {$not: {$in: ", sUDI.RoleActionInfo, "}}}", err)
	}
	clothingNum, err = FindCount(bson.M{"rolename": roleName, "systemrole": 0, "clothingname": bson.M{"$not": bson.M{"$in": sUDI.RoleClothingInfo}}}, clothingInfo)
	if err != nil {
		beego.Error("[rrmhjbg.com/dao/resourceDao.GetRoleTipNum(uid=", uid, "roleName=", roleName, ")] {rolename: ", roleName, ", systemrole: 0, clothingname: {$not: {$in: ", sUDI.RoleClothingInfo, "}}}", err)
	}

	return faceNum + actionNum + clothingNum

}

//2013/08/06 Wangdj 新增：获取当前角色信息
func GetRoleInfo(roleName string, srcInfo *resource.SrcRoleInfo) {
	err := FindOne(bson.M{"rolename": roleName}, srcInfo, roleInfo)
	if err != nil {
		beego.Error("[rrmhjbg.com/dao/resourceDao.GetRoleInfo(roleName=", roleName, ")]", err)
	}
}

//2013/07/17 Wangdj 新增：获取当前用户的已下载资源信息
func GetDownloadInfoByUid(uid string, srcUserDownInfo *resource.SrcUserDownloaded) (err error) {
	err = FindOne(bson.M{"uid": uid}, srcUserDownInfo, userDown)
	if err != nil && err != mgo.ErrNotFound {
		beego.Error("[rrmhjbg.com/dao/resourceDao.GetDownloadInfoByUid(uid=", uid, ")] 获取当前用户的已下载资源信息时出错：", err)
		return err
	}

	return nil
}

//2013/07/23 Wangdj 新增：记录当前用户已经下载过此角色
//2013/07/24 Wangdj 修改：在记录下载信息前，先查询此用户记录是否存在，如果不存在，则新增。
func SaveRoleInUser(roleName, uid string) {
	initDownUser(uid)

	err := Update(userDown, bson.M{"uid": uid}, bson.M{"$addToSet": bson.M{"roleInfo": roleName}})
	if err != nil {
		beego.Error("[rrmhjbg.com/dao/resourceDao.SaveRoleInUser(roleName=", roleName, "uid=", uid, ")] 记录已下载角色信息时出错：", err)
	}
}

//2013/07/24 Wangdj 新增：记录当前用户已经下载过的表情
func SaveRoleFaceInUser(faceNames []string, uid string) {
	initDownUser(uid)
	err := Update(userDown, bson.M{"uid": uid}, bson.M{"$addToSet": bson.M{"roleFaceInfo": bson.M{"$each": faceNames}}})
	if err != nil {
		beego.Error("[rrmhjbg.com/dao/resourceDao.SaveRoleFaceInUser(faceNames=", faceNames, "uid=", uid, ")] 记录当前用户已经下载过的表情时出错：", err)
	}
}

//2013/07/24 Wangdj 新增：记录当前用户已经下载过的动作
func SaveRoleActionInUser(actionNames []string, uid string) {
	initDownUser(uid)
	err := Update(userDown, bson.M{"uid": uid}, bson.M{"$addToSet": bson.M{"roleActionInfo": bson.M{"$each": actionNames}}})
	if err != nil {
		beego.Error("[rrmhjbg.com/dao/resourceDao.SaveRoleActionInUser(actionNames=", actionNames, "uid=", uid, ")] 记录当前用户已经下载过的表情时出错：", err)
	}
}

//2013/07/24 Wangdj 新增：记录当前用户已经下载过的服装
func SaveRoleClothingInUser(clothinsNames []string, uid string) {
	initDownUser(uid)
	err := Update(userDown, bson.M{"uid": uid}, bson.M{"$addToSet": bson.M{"roleClothingInfo": bson.M{"$each": clothinsNames}}})
	if err != nil {
		beego.Error("[rrmhjbg.com/dao/resourceDao.SaveRoleClothingInUser(clothinsNames=", clothinsNames, "uid=", uid, ")] 记录当前用户已经下载过的表情时出错：", err)
	}
}

//2013/07/24 Wangdj 新增：记录当前用户已经下载过的对话框
func SaveDialogInUser(dialogNames []string, uid string) {
	initDownUser(uid)
	err := Update(userDown, bson.M{"uid": uid}, bson.M{"$addToSet": bson.M{"dialogInfo": bson.M{"$each": dialogNames}}})
	if err != nil {
		beego.Error("[rrmhjbg.com/dao/resourceDao.SaveDialogInUser(dialogNames=", dialogNames, "uid=", uid, ")]：", err)
	}
}

//2013/07/24 Wangdj 新增：记录当前用户已经下载过的场景
func SaveSceneInUser(sceneNames []string, uid string) {
	initDownUser(uid)
	err := Update(userDown, bson.M{"uid": uid}, bson.M{"$addToSet": bson.M{"sceneInfo": bson.M{"$each": sceneNames}}})
	if err != nil {
		beego.Error("[rrmhjbg.com/dao/resourceDao.SaveSceneInUser(sceneNames=", sceneNames, "uid=", uid, ")]：", err)
	}
}

func initDownUser(uid string) {
	count, err := FindCount(bson.M{"uid": uid}, userDown)

	if count == 0 || err == mgo.ErrNotFound {
		err = Insert(userDown, bson.M{"uid": uid})
		if err != nil {
			beego.Error("[rrmhjbg.com/dao/resourceDao.initDownUser(uid=", uid, ")]", err)
			return
		}
	}
}

//2013/07/26 Wangdj 新增：清空当前用户的下载记录
func InitUserDownInfo(uid string) {
	result := resource.SrcUserDownloaded{}
	err := FindOne(bson.M{"uid": uid}, &result, userDown)
	if err == nil {
		err = Update(userDown, bson.M{"uid": uid}, bson.M{"$unset": bson.M{"dialogInfo": 1, "roleActionInfo": 1, "roleClothingInfo": 1, "roleFaceInfo": 1, "roleInfo": 1, "sceneInfo": 1}})
		if err != nil {
			beego.Error("[rrmhjbg.com/dao/resourceDao.InitUserDownInfo(uid=", uid, ")]：", err)
		}
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//2013/07/23 Wangdj 新增：获取指定角色的资源 —— 角色
//2013/07/24 Wangdj 修改：具体业务代码改为调用getRoleInfo公共方法
func GetRoleBySystem(roleName string, systemRole int, srcRoleInfo *[]resource.SrcRoleInfo) {
	getRoleInfo(roleName, systemRole, srcRoleInfo, roleInfo)
}

//2013/07/23 Wangdj 新增：获取指定角色的资源 —— 表情
//2013/07/24 Wangdj 修改：具体业务代码改为调用getRoleInfo公共方法
func GetRoleFaceBySystem(roleName string, systemRole int, srcRoleFaceInfo *[]resource.SrcRoleFaceInfo) {
	getRoleInfo(roleName, systemRole, srcRoleFaceInfo, faceInfo)
}

//2013/07/23 Wangdj 新增：获取指定角色的资源 —— 动作与衣服
//2013/07/24 Wangdj 修改：具体业务代码改为调用getRoleInfo公共方法
func GetRoleActionClothingBySystem(roleName string, systemRole int, srcRoleActionInfo *[]resource.SrcRoleActionInfo) {
	getRoleInfo(roleName, systemRole, srcRoleActionInfo, actionInfo)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//2013/07/25 Wangdj 新增：查找指定表情
func GetOneRoleFaceByKey(faceName string, srcRoleFaceInfo *resource.SrcRoleFaceInfo) (found bool) {
	err := FindOne(bson.M{"facename": faceName}, srcRoleFaceInfo, faceInfo)
	if err != nil {
		beego.Error("[rrmhjbg.com/dao/resourceDao.GetOneRoleFaceByKey(faceName=", faceName, ")]：", err)
		return false
	}

	return true
}

//2013/07/25 Wangdj 新增：查找指定动作
func GetOneRoleActionByKey(actionName string, srcRoleActionInfo *resource.SrcRoleActionInfo) (found bool) {
	err := FindOne(bson.M{"actionname": actionName}, srcRoleActionInfo, actionInfo)
	if err != nil {
		beego.Error("[rrmhjbg.com/dao/resourceDao.GetOneRoleActionByKey(actionName=", actionName, ")]：", err)
		return false
	}

	return true
}

//2013/07/25 Wangdj 新增：查找指定服装
func GetOneRoleClothingByKey(clothingName string, srcRoleClothingInfo *resource.SrcRoleClothingInfo, srcRoleActionInfo *[]resource.SrcRoleActionInfo) (found bool) {
	err := FindOne(bson.M{"clothingname": clothingName}, srcRoleClothingInfo, clothingInfo)
	if err != nil {
		beego.Error("[rrmhjbg.com/dao/resourceDao.GetOneRoleClothingByKey(clothingName=", clothingName, ")]：", err)
		return false
	}

	_, err = FindList(bson.M{"clothing.clothinggroup": bson.M{"$all": []string{srcRoleClothingInfo.ClothingGroup}}, "rolename": srcRoleClothingInfo.RoleName}, srcRoleActionInfo, actionInfo, 0, 1000, "posttime")

	if err != nil {
		beego.Error("[rrmhjbg.com/dao/resourceDao.GetOneRoleClothingByKey(clothingName=", clothingName, ")]：", err)
		return false
	}

	return true
}

//2013/07/25 Wangdj 新增：查找指定对话框
func GetOneDialogByKey(dialogName string, srcDialogInfo *resource.SrcDialogInfo) (found bool) {
	err := FindOne(bson.M{"dialogname": dialogName}, srcDialogInfo, dialogInfo)
	if err != nil {
		beego.Error("[rrmhjbg.com/dao/resourceDao.GetOneDialogByKey(dialogName=", dialogName, ")]：", err)
		return false
	}

	return true
}

//2013/07/25 Wangdj 新增：查找指定场景
func GetOneSceneByKey(sceneName string, srcSceneInfo *resource.SrcSceneInfo) (found bool) {
	err := FindOne(bson.M{"scenename": sceneName}, srcSceneInfo, sceneInfo)
	if err != nil {
		beego.Error("[rrmhjbg.com/dao/resourceDao.GetOneSceneByKey(sceneName=", sceneName, ")]：", err)
		return false
	}

	return true
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////////////////
//2013/07/25 Wangdj 新增：用于将excel中的资源数据导入到数据库中。                          ///////////////
///////////////////////////////////////////////////////////////////////////////////////////////////
func InsertRole(srcRoleInfo *resource.SrcRoleInfo) {
	srcRoleInfo.Id = bson.NewObjectId().Hex()
	srcRoleInfo.PostTime = time.Now()
	srcRoleInfo.Iflag = NormalCode

	insertSrc(roleInfo, "rolename", srcRoleInfo.RoleName, srcRoleInfo)
}

func InsertRoleFace(srcRoleFaceInfo *resource.SrcRoleFaceInfo) {
	srcRoleFaceInfo.Id = bson.NewObjectId().Hex()
	srcRoleFaceInfo.PostTime = time.Now()
	srcRoleFaceInfo.Iflag = NormalCode

	insertSrc(faceInfo, "facename", srcRoleFaceInfo.FaceName, srcRoleFaceInfo)
}

func InsertDialog(srcInfo *resource.SrcDialogInfo) {
	srcInfo.Id = bson.NewObjectId().Hex()
	srcInfo.PostTime = time.Now()
	srcInfo.Iflag = NormalCode

	insertSrc(dialogInfo, "dialogname", srcInfo.DialogName, srcInfo)
}

func InsertScene(srcInfo *resource.SrcSceneInfo) {
	srcInfo.Id = bson.NewObjectId().Hex()
	srcInfo.PostTime = time.Now()
	srcInfo.Iflag = NormalCode

	insertSrc(sceneInfo, "scenename", srcInfo.SceneName, srcInfo)
}

func InsertRoleAction(srcInfo *resource.SrcRoleActionInfo) {
	temp := resource.SrcRoleActionInfo{}
	err := FindOne(bson.M{"itempicname": srcInfo.ItemPicName}, &temp, actionInfo)
	if err == mgo.ErrNotFound {
		srcInfo.Id = bson.NewObjectId().Hex()
		srcInfo.PostTime = time.Now()
		srcInfo.Iflag = NormalCode
		srcInfo.ActionName = srcInfo.Id

		err = Insert(actionInfo, srcInfo)
		if err != nil {
			beego.Error(err)
		}

		return
	} else if err != nil {
		beego.Error(err)
		return
	}

	err = Update(actionInfo, bson.M{"itempicname": srcInfo.ItemPicName}, bson.M{"$push": bson.M{"clothing": srcInfo.Clothing[0]}})
	if err != nil {
		beego.Error(err)
	}
}

func InsertRoleClothing(srcInfo *resource.SrcRoleClothingInfo) {
	temp := resource.SrcRoleActionInfo{}
	err := FindOne(bson.M{"clothinggroup": srcInfo.ClothingGroup, "rolename": srcInfo.RoleName}, &temp, clothingInfo)
	if err == mgo.ErrNotFound {
		srcInfo.Id = bson.NewObjectId().Hex()
		srcInfo.PostTime = time.Now()
		srcInfo.Iflag = NormalCode

		err = Insert(clothingInfo, srcInfo)
		if err != nil {
			beego.Error(err)
		}
	} else if err != nil {
		beego.Error(err)
	}
}

func insertSrc(collection, seletorKey, selectorVal string, data interface{}) {
	err := Remove(collection, bson.M{seletorKey: selectorVal})
	if err != nil && err != mgo.ErrNotFound {
		beego.Error(err)
		return
	}

	err = Insert(collection, data)
	if err != nil {
		beego.Error(err)
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//2013/07/24 Wangdj 新增：获取角色及相关资源的公共方法（角色，角色下的表情、动作与衣服）
func getRoleInfo(roleName string, isSystem int, result interface{}, collectionName string) {
	query := bson.M{"rolename": roleName, "systemrole": isSystem}
	if isSystem == -1 {
		query = bson.M{"rolename": roleName}
	}
	_, err := FindList(query, result, collectionName, 0, 1000, "-sort")
	if err != nil && err != mgo.ErrNotFound {
		var srcName string
		var systemRole string

		if isSystem != 1 {
			systemRole = "非"
		}
		switch collectionName {
		case roleInfo:
			srcName = "角色"
		case faceInfo:
			srcName = "角色下表情"
		case actionInfo:
			srcName = "角色下动作与衣服"
		}

		beego.Error("[rrmhjbg.com/dao/resourceDao.GetRoleBySystem(roleName=", roleName, ")] 获取指定"+srcName+"的"+systemRole+"系统（默认）资源时出错：", err)
		return
	}
}

//2013/07/18 Wangdj 新增：将需要分页获取的资源列表整理成一个公共方法，供不同类型资源调用
func showResourceInfoByPage(pageIndex, pageSize int, tableName, roleName string, result interface{}) (count int, err error) {

	query := bson.M{"iflag": NormalCode}
	if roleName != "" {
		query = bson.M{"iflag": NormalCode, "rolename": roleName}
	}

	count, err = FindList(query, result, tableName, (pageIndex-1)*pageSize, pageSize, "-sort")

	if err != nil && err != mgo.ErrNotFound {
		var funcName, errInfo string
		switch tableName {
		case roleInfo:
			funcName = "ShowRoleInfoByPage"
			errInfo = "根据分页获取角色列表"

		case dialogInfo:
			funcName = "ShowDialogInfoByPage"
			errInfo = "根据分页获取对话框列表"

		case sceneInfo:
			funcName = "ShowSceneInfoByPage"
			errInfo = "根据分页获取场景列表"

		case faceInfo:
			funcName = "ShowRoleFaceInfoByPage"
			errInfo = "根据分页获取相应角色表情列表"

		case actionInfo:
			funcName = "ShowRoleActionInfoByPage"
			errInfo = "根据分页获取相应角色动作列表"

		case clothingInfo:
			funcName = "ShowRoleClothingInfoByPage"
			errInfo = "根据分页获取相应角色服装列表"
		}

		beego.Error("[rrmhjbg.com/dao/resourceDao."+funcName+"(pageIndex=", pageIndex, ",pageSize=", pageSize, ")] "+errInfo+"时出错：", err)
		return 0, err
	}

	return count, nil
}
