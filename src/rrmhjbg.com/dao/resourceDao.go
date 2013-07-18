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
	"rrmhjbg.com/conf"
	"rrmhjbg.com/models/resource"
)

const (
	roleInfo   = "srcRoleInfo"
	dialogInfo = "srcDialogInfo"
	sceneInfo  = "srcSceneInfo"

	NormalCode = 0
	DelCode    = -1
)

//2013/07/17 Wangdj 新增：根据分页获取角色列表
func ShowRoleInfoByPage(pageIndex, pageSize int, srcRoleInfo *[]resource.SrcRoleInfo) (count int, err error) {
	return showResourceInfoByPage(pageIndex, pageSize, roleInfo, srcRoleInfo)
}

//2013/07/18 Wangdj 新增：根据分页获取对话框列表
func ShowDialogInfoByPage(pageIndex, pageSize int, srcDialogInfo *[]resource.SrcDialogInfo) (count int, err error) {
	return showResourceInfoByPage(pageIndex, pageSize, dialogInfo, srcDialogInfo)
}

//2013/07/18 Wangdj 新增：根据分页获取场景列表
func ShowSceneInfoByPage(pageIndex, pageSize int, srcSceneInfo *[]resource.SrcSceneInfo) (count int, err error) {
	return showResourceInfoByPage(pageIndex, pageSize, sceneInfo, srcSceneInfo)
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

//2013/07/18 Wangdj 新增：将需要分页获取的资源列表整理成一个公共方法，供不同类型资源调用
func showResourceInfoByPage(pageIndex, pageSize int, tableName string, result interface{}) (count int, err error) {

	count, err = FindList(bson.M{"iflag": NormalCode}, result, tableName, (pageIndex-1)*pageSize, pageSize, "-postTime")

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
		}

		beego.Error("[rrmhjbg.com/dao/resourceDao."+funcName+"(pageIndex=", pageIndex, ",pageSize=", pageSize, ")] "+errInfo+"时出错：", err)
		return 0, err
	}

	return count, nil
}

//2013/07/17 Wangdj 新增：将资源库设置为当前数据库
func init() {
	DBName = conf.ResourceDBName
	beego.Debug("init test")
}
