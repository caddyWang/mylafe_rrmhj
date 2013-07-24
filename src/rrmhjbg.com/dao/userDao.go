package dao

/************************************************************************************
//
// Desc		:	用户数据Dao
// Records	:	2013-06-18	Wangdj	新建文件；增加函数"InitUserInfoBySinaWeibo"
//
************************************************************************************/

import (
	"github.com/astaxie/beego"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"rrmhjbg.com/models"
	"time"
)

const userInfo = "userInfo"

func InitUserInfoBySinaWeibo(user *models.UserInfo, platformName string) (uid string) {
	var err error
	var fieldName string
	var pushSocialUser map[string]interface{}

	switch platformName {
	case models.SinaWeibo:
		fieldName = "sinaweibo"
		pushSocialUser, err = findSocialUser(fieldName, user.SinaWeibo[0], user)

	case models.TencWeibo:
		fieldName = "tencweibo"
		pushSocialUser, err = findSocialUser(fieldName, user.TencWeibo[0], user)

	case models.QQZone:
		fieldName = "qqzone"
		pushSocialUser, err = findSocialUser(fieldName, user.QQZone[0], user)

	case models.RenRenSNS:
		fieldName = "renrensns"
		pushSocialUser, err = findSocialUser(fieldName, user.RenRenSNS[0], user)
	}

	if err == mgo.ErrNotFound {
		if user.Id == "-1" {
			user.Id = bson.NewObjectId().Hex()
			user.CreateTime = time.Now()

			err = InsertDB(userInfo, DBNameForUser, user)
			if err != nil {
				beego.Error("新增用户信息时出错：", *user, err)
			}
		} else {
			err = UpdateDB(userInfo, DBNameForUser, bson.M{"_id": user.Id}, bson.M{"$push": bson.M{fieldName: pushSocialUser}})
		}
	} else if err != nil {
		beego.Error("查询用户数据出错：platformName=", platformName, err)
	}

	return user.Id
}

func GetAllUsers(pageIndex, pageSize int, sort string, user *[]models.UserInfo) (icount int, err error) {
	if sort == "" {
		sort = "-createtime"
	}
	icount, err = FindListDB(bson.M{}, user, userInfo, DBNameForUser, (pageIndex-1)*pageSize, pageSize, sort)
	if err != nil {
		beego.Error("查询用户分页数据出错：pageindex=", pageIndex, " pagesize=", pageSize, err)
	}
	return
}

func UpateUser(user *models.UserInfo) {
	UpdateDB(userInfo, DBNameForUser, bson.M{"_id": user.Id}, user)
}

func findSocialUser(platform string, socialUser models.SocialUserInfo, user *models.UserInfo) (pushSocialUser map[string]interface{}, err error) {
	err = FindOneDB(bson.M{platform: bson.M{"$elemMatch": bson.M{"uid": socialUser.Uid}}}, user, userInfo, DBNameForUser)
	user.UserName, user.ProfileImg = socialUser.UserName, socialUser.ProfileImg
	pushSocialUser = bson.M{"uid": socialUser.Uid, "username": socialUser.UserName, "profileimg": socialUser.ProfileImg, "profileurl": socialUser.ProfileUrl}

	beego.Debug("{", platform, ":{$elemMatch:{uid:", socialUser.Uid, "}}}")
	beego.Debug(user)
	return
}
