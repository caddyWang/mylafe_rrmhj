package business

/************************************************************************************
//
// Desc		:	与会员相关的业务功能
// Records	:	2013-06-14	Wangdj	新建文件；增加函数"CheckLogin"
//
************************************************************************************/

import (
	"github.com/astaxie/beego"
	"rrmhj.com/dao"
	"rrmhj.com/models"
)

type GetSession func(key interface{}) interface{}
type SetSession func(name interface{}, value interface{})

//验证用户是否登录
func CheckLogin(gs GetSession) bool {
	uid := gs("uid")

	if uid == nil {
		return false
	}

	if uid == "" {
		return false
	}

	return true
}

//获取已登录用户的id,username,profileimg(wangdj 2013-06-19)
func GetSessinUserBase(gs GetSession) (user models.UserBase) {
	var uidStr, unameStr, uprofileimgStr string
	uid, uname, uprofileimg := gs("uid"), gs("uname"), gs("uprofileimg")
	if uid != nil {
		uidStr = uid.(string)
	}
	if uname != nil {
		unameStr = uname.(string)
	}
	if uprofileimg != nil {
		uprofileimgStr = uprofileimg.(string)
	}

	user = models.UserBase{uidStr, unameStr, uprofileimgStr}
	return
}

func SinaLoginProcess(tkRST *SinaWeiboOauth2AccesstokenResult, userRST *SinaWeiboUserShowResult, setSess SetSession) {
	if userRST.Id <= 0 {
		beego.Error("获取新浪用户信息数据出错，不能正常登录！")
		return
	}

	user := models.UserInfo{Id: "-1", Gender: userRST.Gender, Province: userRST.Province, City: userRST.City, Location: userRST.Location}
	user.SinaWeibo = models.SinaWeiboUserInfo{userRST.Id, userRST.Screen_name, userRST.Profile_image_url, userRST.Avatar_large, userRST.Profile_url, userRST.Description}

	userId := dao.InitUserInfoBySinaWeibo(&user)

	setSess("sina_access_token", tkRST.Access_token)
	setSess("sina_id", userRST.Id)
	setSess("uid", userId)
	setSess("uname", user.UserName)
	setSess("uprofileimg", user.ProfileImg)
}
