package dao

/************************************************************************************
//
// Desc		:	用户数据Dao
// Records	:	2013-06-18	Wangdj	新建文件；增加函数"InitUserInfoBySinaWeibo"
//
************************************************************************************/

import (
	"github.com/astaxie/beego"
	"labix.org/v2/mgo/bson"
	"rrmhj.com/models"
	"time"
)

const userInfo = "userInfo"

func InitUserInfoBySinaWeibo(user *models.UserInfo) (uid string) {

	err := FindOne(bson.M{"sinaweibo.snuid": user.SinaWeibo.SnUid}, user, userInfo)
	if err != nil {
		if err.Error() != "not found" {
			beego.Error("查询用户数据出错：", bson.M{"sinaweibo.snuid": user.SinaWeibo.SnUid}, err)
		}
		beego.Debug("查询用户数据：", bson.M{"sinaweibo.snuid": user.SinaWeibo.SnUid}, err)
	}

	if user.Id == "-1" {
		user.Id = bson.NewObjectId().Hex()
		user.UserName = user.SinaWeibo.SnUserName
		user.ProfileImg = user.SinaWeibo.SnProfileImg
		user.ProfileLargeImg = user.SinaWeibo.SnAvaterLarge
		user.CreateTime = time.Now()

		err = Insert(userInfo, user)
		if err != nil {
			beego.Error("新增用户信息时出错：", *user, err)
		}
	}

	return user.Id
}
