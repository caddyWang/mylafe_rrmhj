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
	"strings"
)

//如果用户没用设置过头像，则获取默认头像（Wangdj 2013-06-07	）
func DefaultHeadImg(headImg string) string {
	if strings.Trim(headImg, " ") == "" {
		return beego.AppConfig.String("defheadimg")
	}

	return headImg
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
