package business

/************************************************************************************
//
// Desc		:	页面模板函数
// Records	:	2013-06-09	Wangdj	新建文件；增加函数"DefaultHeadImg"
//				2013-06-14	Wangdj	新建文件；增加函数"LoginDisplay"、"LogoutDisplay"
//
************************************************************************************/

import (
	"github.com/astaxie/beego"
	"html/template"
	"rrmhj.com/conf"
	"strings"
)

//如果用户没用设置过头像，则获取默认头像（Wangdj 2013-06-07	）
func DefaultHeadImg(headImg string) string {

	beego.Debug("头像地址：", headImg)
	beego.Debug("是否为站外地址：", strings.Contains(headImg, "http:"))

	if strings.Trim(headImg, " ") == "" {
		headImg = conf.DefProfileImg
	}

	if strings.Contains(headImg, "http:") {
		return headImg
	}

	return conf.StaticFileURL + "/" + headImg
}

//当用户登录时，显示html控件
func LoginDisplay(islogin bool) template.HTMLAttr {

	//beego.Info(islogin)

	if islogin {
		return template.HTMLAttr("")
	}

	return template.HTMLAttr("style='display:none'")
}

//当用户未登录时，显示html控件
func LogoutDisplay(islogin bool) template.HTMLAttr {
	return LoginDisplay(!islogin)
}

// 2013/07/11 Wangdj 新增：判断页面上显示的作品是否是当前用户已经收藏的
func DisplayLike(userId, proId string) template.HTML {
	if IsLike(userId, proId) == "" {
		return template.HTML("<div class='star-full'></div>已收藏")
	}

	return template.HTML("<div class='star'></div>收藏")
}

// 2013/07/11 Wangdj 新增：判断页面上显示的作品是否是当前用户已经收藏的
func IsLike(userId, proId string) string {
	if userId == "" || proId == "" {
		return "unlike"
	}

	pros := GetUserLikeProduct(userId)

	beego.Debug("Like:userId=", userId, ",proId=", proId)
	beego.Debug("Like:pros=", pros)

	for _, p := range pros {
		if p == proId {
			return ""
		}
	}

	return "unlike"
}
