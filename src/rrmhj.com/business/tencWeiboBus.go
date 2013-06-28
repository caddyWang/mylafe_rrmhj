package business

import (
	"rrmhj.com/conf"
)

type TencWeiboOauth2AccesstokenResult struct {
	Access_token  string //用于调用access_token，接口获取授权后的access token。
	Expires_in    string //access_token的生命周期，单位是秒数。
	Refresh_token string
}

type TencWeiboUserShowResult struct {
	Data struct {
		Openid        string //字符串型的用户UID
		Name          string //用户昵称
		Nick          string //友好显示名称
		Province_code string //用户所在省级ID
		City_code     string //用户所在城市ID
		Location      string //用户所在地
		Introduction  string //用户个人描述
		Homepage      string //用户博客地址
		Head          string //用户头像地址，50×50像素
		Sex           string //性别，m：男、f：女、n：未知
	}
}

//从配置文件中获取的appKey和appSecret
func GetTencWeiboKeys() (appKey, appSecret string) {
	return conf.TencClientId, conf.TencClientSecret
}

//从配置文件中读取OAuth登录验证所需要的信息
func GetTencWeiboOAuthParams() (addr, reduri, grantType string) {
	return conf.TencOauth2AccesstokenAddr, conf.TencRedirectUri, conf.TencGrantType
}

//从配置文件中读取获取用户信息所需要的信息
func GetTencWeiboAPIUrls() (userShow string) {
	return conf.TencUserShowAddr
}
