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
	"rrmhjbg.com/models"
	"time"
)

const userInfo = "userInfo"

func InitUserInfoBySinaWeibo(user *models.UserInfo, platformName string) (uid string) {
	var err error
	switch platformName {
	case models.SinaWeibo:
		err = FindOne(bson.M{"SinaWeibo": bson.M{"$elemMatch": bson.M{"uid": user.SinaWeibo[0].Uid}}}, user, userInfo)
		user.UserName, user.ProfileImg = user.SinaWeibo[0].UserName, user.SinaWeibo[0].ProfileImg

	case models.TencWeibo:
		err = FindOne(bson.M{"TencWeibo": bson.M{"$elemMatch": bson.M{"uid": user.TencWeibo[0].Uid}}}, user, userInfo)
		user.UserName, user.ProfileImg = user.TencWeibo[0].UserName, user.TencWeibo[0].ProfileImg

	case models.QQZone:
		err = FindOne(bson.M{"QQZone": bson.M{"$elemMatch": bson.M{"uid": user.QQZone[0].Uid}}}, user, userInfo)
		user.UserName, user.ProfileImg = user.QQZone[0].UserName, user.QQZone[0].ProfileImg

	case models.RenRenSNS:
		err = FindOne(bson.M{"RenRenSNS": bson.M{"$elemMatch": bson.M{"uid": user.RenRenSNS[0].Uid}}}, user, userInfo)
		user.UserName, user.ProfileImg = user.RenRenSNS[0].UserName, user.RenRenSNS[0].ProfileImg
	}

	if err != nil {
		if err.Error() != "not found" {
			beego.Error("查询用户数据出错：platformName=", platformName, err)
		}
		beego.Debug("查询用户数据：platformName=", platformName, err)
	}

	if user.Id == "-1" {
		user.Id = bson.NewObjectId().Hex()
		user.CreateTime = time.Now()

		err = Insert(userInfo, user)
		if err != nil {
			beego.Error("新增用户信息时出错：", *user, err)
		}
	}

	return user.Id
}
