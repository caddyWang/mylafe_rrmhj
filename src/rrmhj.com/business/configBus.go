package business

import (
	"github.com/astaxie/beego"
)

type SinaWeiboOauth2AccesstokenResult struct {
	Access_token string //用于调用access_token，接口获取授权后的access token。
	Expires_in   int    //access_token的生命周期，单位是秒数。
	Remind_in    string //access_token的生命周期（该参数即将废弃，开发者请使用expires_in）。
	Uid          string //当前授权用户的UID。
}

type SinaWeiboUserShowResult struct {
	Id                int64  //用户UID
	Idstr             string //字符串型的用户UID
	Screen_name       string //用户昵称
	Name              string //友好显示名称
	Province          string //用户所在省级ID
	City              string //用户所在城市ID
	Location          string //用户所在地
	Description       string //用户个人描述
	Url               string //用户博客地址
	Profile_url       string //简地址，不含weibo.com前缀，如：u/3125160187
	Profile_image_url string //用户头像地址，50×50像素
	Gender            string //性别，m：男、f：女、n：未知
	Avatar_large      string //用户大头像地址
}

func GetSinaWeiboKeys() (appKey, appSecret string) {
	appKey = beego.AppConfig.String("sina_client_id")
	appSecret = beego.AppConfig.String("sina_client_secret")

	return
}

func GetSinaWeiboOAuthParams() (addr, reduri, grantType string) {
	addr = beego.AppConfig.String("sina_oauth2_accesstoken_addr")
	reduri = beego.AppConfig.String("sina_redirect_uri")
	grantType = beego.AppConfig.String("sina_grant_type")

	return
}

func GetSinaWeiboAPIUrls() (userShow string) {
	userShow = "https://api.weibo.com/2/users/show.json" //根据用户ID获取用户信息

	return
}
