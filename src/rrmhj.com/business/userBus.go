package business

/************************************************************************************
//
// Desc		:	与会员相关的业务功能
//
************************************************************************************/

import (
	"github.com/astaxie/beego"
	"rrmhj.com/dao"
	"rrmhj.com/models"
	"strconv"
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

// 2013/06/26 Wangdj 调整：用户数据结构发生变化。一个用户帐号，会绑定多个相同平台的社交帐号，比如一个用户可能会有多个新浪帐号，所以sinaweibo、tencweibo、qqzone、renrensns字段变成数组类型
func SinaLoginProcess(tkRST *SinaWeiboOauth2AccesstokenResult, userRST *SinaWeiboUserShowResult, setSess SetSession) {
	if userRST.Id <= 0 {
		beego.Error("获取新浪用户信息数据出错，不能正常登录！")
		return
	}

	user := models.UserInfo{Id: "-1", Gender: userRST.Gender, Province: userRST.Province, City: userRST.City, Location: userRST.Location}
	socialUser := models.SocialUserInfo{strconv.FormatInt(userRST.Id, 10), userRST.Screen_name, userRST.Profile_image_url, userRST.Profile_url, userRST.Gender, userRST.Province, userRST.City, userRST.Location, userRST.Avatar_large, userRST.Description}
	user.SinaWeibo = []models.SocialUserInfo{socialUser}

	userId := dao.InitUserInfoBySocialUser(&user, models.SinaWeibo)

	setSess("sina_access_token", tkRST.Access_token)
	setSess("sina_id", userRST.Id)
	setSess("uid", userId)
	setSess("uname", user.UserName)
	setSess("uprofileimg", user.ProfileImg)
}
