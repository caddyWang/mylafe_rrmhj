package dao

/************************************************************************************
//
// Desc		:	用户数据Dao
//
************************************************************************************/

import (
	"github.com/astaxie/beego"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"rrmhj.com/models"
	"time"
)

const userInfo = "userInfo"

func InitUserInfoBySocialUser(user *models.UserInfo, platformName string) (uid string) {
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

			err = Insert(userInfo, user)
			if err != nil {
				beego.Error("新增用户信息时出错：", *user, err)
			}
		} else {
			err = Update(userInfo, bson.M{"_id": user.Id}, bson.M{"$push": bson.M{fieldName: pushSocialUser}})
		}
	} else if err != nil {
		beego.Error("查询用户数据出错：platformName=", platformName, err)
	}

	return user.Id
}

// 2013/07/10 Wangdj 新增：用户收藏作品功能
func SaveUserLikeProduct(proId, userId string) (err error) {
	change := bson.M{"$push": bson.M{"likepro": proId}}

	err = Update(userInfo, bson.M{"_id": userId}, change)
	if err != nil {
		beego.Error("添加用户收藏商品时出错：proid=", proId, ", userId=", userId, err)
	}

	return
}

// 2013/07/11 Wangdj 新增：查找当前用户已经收藏的作品
func GetUserLikeProduct(userId string) []string {
	user := models.UserInfo{}
	err := FindOne(bson.M{"_id": userId}, &user, userInfo)
	if err != nil {
		beego.Error("查找当前用户已经收藏的作品出错：userId=", userId, err)
	}

	return user.LikePro
}

func findSocialUser(platform string, socialUser models.SocialUserInfo, user *models.UserInfo) (pushSocialUser map[string]interface{}, err error) {
	err = FindOne(bson.M{platform: bson.M{"$elemMatch": bson.M{"uid": socialUser.Uid}}}, user, userInfo)
	user.UserName, user.ProfileImg = socialUser.UserName, socialUser.ProfileImg
	pushSocialUser = bson.M{"uid": socialUser.Uid, "username": socialUser.UserName, "profileimg": socialUser.ProfileImg, "profileurl": socialUser.ProfileUrl}

	beego.Debug("{", platform, ":{$elemMatch:{uid:", socialUser.Uid, "}}}")
	beego.Debug(user)
	return
}
