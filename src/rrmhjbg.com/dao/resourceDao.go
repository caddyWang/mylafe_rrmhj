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

//2013/07/17 Wangdj 新增：获取当前用户的已下载资源信息
func GetDownloadInfoByUid(uid string, srcUserDownInfo *resource.SrcUserDownloaded) (err error) {
	err = FindOne(bson.M{"uid": uid}, srcUserDownInfo, roleInfo)
	if err != nil && err != mgo.ErrNotFound {
		beego.Error("[rrmhjbg.com/dao/resourceDao.GetDownloadInfoByUid(uid=", uid, ")] 获取当前用户的已下载资源信息时出错：", err)
		return err
	}

	return nil
}

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

//2013/07/23 Wangdj 新增：记录当前用户已经下载过此角色
//2013/07/24 Wangdj 修改：在记录下载信息前，先查询此用户记录是否存在，如果不存在，则新增。
func SaveRoleInUser(roleName, uid string) {
	result := resource.SrcUserDownloaded{}
	err := FindOne(bson.M{"uid": uid}, &result, userDown)

	if err == mgo.ErrNotFound {
		err = Insert(userDown, bson.M{"uid": uid})
		if err != nil {
			beego.Error("[rrmhjbg.com/dao/resourceDao.SaveRoleInUser(roleName=", roleName, "uid=", uid, ")] 新增已下载角色信息时出错：", err)
			return
		}
	} else if err != nil {
		beego.Error("[rrmhjbg.com/dao/resourceDao.SaveRoleInUser(roleName=", roleName, "uid=", uid, ")] 查询已下载角色信息时出错：", err)
		return
	}

	err = Update(userDown, bson.M{"uid": uid}, bson.M{"$addToSet": bson.M{"roleInfo": roleName}})
	if err != nil {
		beego.Error("[rrmhjbg.com/dao/resourceDao.SaveRoleInUser(roleName=", roleName, "uid=", uid, ")] 记录已下载角色信息时出错：", err)
	}
}

//2013/07/24 Wangdj 新增：获取角色及相关资源的公共方法（角色，角色下的表情、动作与衣服）
func getRoleInfo(roleName string, isSystem int, result interface{}, collectionName string) {
	_, err := FindList(bson.M{"rolename": roleName, "systemrole": isSystem}, result, collectionName, 0, 1000, "posttime")
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

	count, err = FindList(query, result, tableName, (pageIndex-1)*pageSize, pageSize, "-postTime")

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
