package business

import (
	"github.com/astaxie/beego"
	"rrmhj.com/conf"
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

//从配置文件中获取新浪的appKey和appSecret
func GetSinaWeiboKeys() (appKey, appSecret string) {
	return conf.SinaClientId, conf.SinaClientSecret
}

//从配置文件中读取新浪OAuth登录验证所需要的信息
func GetSinaWeiboOAuthParams() (addr, reduri, grantType string) {
	return conf.SinaOauth2AccesstokenAddr, conf.SinaRedirectUri, conf.SinaGrantType
}

//从配置文件中读取获取新浪用户信息所需要的信息
func GetSinaWeiboAPIUrls() (userShow string) {
	beego.Debug("sina_user_show_addr=", beego.AppConfig.String("sina_user_show_addr"))
	beego.Debug("SinaUserShowAddr=", conf.SinaUserShowAddr)
	return conf.SinaUserShowAddr
}
