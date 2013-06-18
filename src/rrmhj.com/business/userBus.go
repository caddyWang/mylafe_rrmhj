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
}
