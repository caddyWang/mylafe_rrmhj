package main

import (
	"github.com/astaxie/beego"
	"rrmhj.com/controllers"
	"strings"
)

func main() {
	beego.Router("/", &controllers.MainController{})

	beego.AddFuncMap("fmtHeadImg", DefaultHeadImg)

	beego.Run()
}

//////////////////////////////////////////////////////////////////////////////
//模板函数
//////////////////////////////////////////////////////////////////////////////

//如果用户没用设置过头像，则获取默认头像（Wangdj 2013-06-07	）
func DefaultHeadImg(headImg string) string {
	if strings.Trim(headImg, " ") == "" {
		return "img/defuser.png"
	}

	return headImg
}
