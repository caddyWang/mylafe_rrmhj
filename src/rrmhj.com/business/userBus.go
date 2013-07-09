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
	setSess("open_platform", "新浪微博")
	setSess("uid", userId)
	setSess("uname", user.UserName)
	setSess("uprofileimg", user.ProfileImg)
}

// 2013/06/28 Wangdj 新增：通过腾讯微博开放平台获取登录用户帐户信息，并绑定到人人漫画家平台的用户信息中
func TencLoginProcess(tkRST *TencWeiboOauth2AccesstokenResult, userRST *TencWeiboUserShowResult, setSess SetSession) {
	if userRST.Data.Openid == "" {
		beego.Error("获取腾讯用户信息数据出错，不能正常登录！")
		return
	}

	user := models.UserInfo{Id: "-1", Gender: userRST.Data.Sex, Province: userRST.Data.Province_code, City: userRST.Data.City_code, Location: userRST.Data.Location}
	socialUser := models.SocialUserInfo{userRST.Data.Openid, userRST.Data.Name, userRST.Data.Head, userRST.Data.Homepage, userRST.Data.Sex, userRST.Data.Province_code, userRST.Data.City_code, userRST.Data.Location, "", userRST.Data.Introduction}
	user.TencWeibo = []models.SocialUserInfo{socialUser}

	userId := dao.InitUserInfoBySocialUser(&user, models.TencWeibo)

	setSess("tenc_access_token", tkRST.Access_token)
	setSess("tenc_id", userRST.Data.Openid)
	setSess("open_platform", "腾讯微博")
	setSess("uid", userId)
	setSess("uname", user.UserName)
	setSess("uprofileimg", user.ProfileImg)
}

// 2013/07/09 Wangdj 新增：获取当前站点登录的用户信息
func LoginedUserInfo(ctx *beego.Controller) {
	ctx.Data["UserName"] = ctx.GetSession("uname")
	ctx.Data["Uid"] = ctx.GetSession("uid")
	ctx.Data["Platform"] = ctx.GetSession("open_platform")
}
