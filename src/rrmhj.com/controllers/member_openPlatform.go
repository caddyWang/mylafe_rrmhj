package controllers

/************************************************************************************
//
// Desc		:   涉及第三方平台相关的用户操作
// Records	:	2013-06-15	Wangdj	新建文件；增加函数"requestSinaAPI"
//
************************************************************************************/

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"net/url"
	"rrmhj.com/business"
	"strings"
)

type SinaLoginController struct {
	beego.Controller
}

// 新浪微博登录操作
func (this *SinaLoginController) Get() {

	this.TplNames = "member/sinalogin.tpl"
	code := this.GetString("code")

	key, secret := business.GetSinaWeiboKeys()
	addr, ruri, gtype := business.GetSinaWeiboOAuthParams()

	//获取新浪微博登录成功后的"access_token"值
	data := url.Values{"client_id": {key}, "client_secret": {secret}, "grant_type": {gtype}, "code": {code}, "redirect_uri": {ruri}}
	beego.Debug("请求新浪微博access_token时所需要的参数：", data)

	rlt := business.SinaWeiboOauth2AccesstokenResult{}
	if err := requestSinaAPI(addr, "POST", "access_token", data, &rlt); err != nil {
		return
	}

	//通过access_token和uid，获取用户详细信息
	addr = business.GetSinaWeiboAPIUrls() + "?access_token=" + rlt.Access_token + "&uid=" + rlt.Uid
	user := business.SinaWeiboUserShowResult{}
	if err := requestSinaAPI(addr, "GET", "用户信息", nil, &user); err != nil {
		return
	}

	business.SinaLoginProcess(&rlt, &user, this.SetSession)
	beego.Debug("Session['uid']", this.GetSession("uid"))

}

//请求新浪微博OpenAPI，获取用户相关信息，并把返回的json数据解析到对应的数据结构中
//reqUrl ： 新浪开发平台相应api功能请求地址
//reqType : 请求方式(post, get)
//reqData : 请求数据，主要用于post中，get通过url中的拼接参数解决
//resultStruct : 返回数据需要解析到相应数据结构指针
func requestSinaAPI(reqUrl, reqType, reqInfo string, reqData map[string][]string, resultStruct interface{}) (err error) {

	var resp *http.Response
	var result []byte

	switch strings.ToUpper(reqType) {
	case "POST":
		resp, err = http.PostForm(reqUrl, reqData)
	default:
		resp, err = http.Get(reqUrl)
	}

	if err != nil {
		beego.Error("请求新浪微博<"+reqInfo+">出错：", err)
		return
	}

	result, err = ioutil.ReadAll(resp.Body) //取出主体的内容
	defer resp.Body.Close()
	if err != nil {
		beego.Error("读取新浪微博<"+reqInfo+">返回值出错:", err)
		return
	}
	beego.Debug("新浪微博返回<"+reqInfo+">的相关值：", string(result[:]))

	json.Unmarshal(result, resultStruct)
	beego.Debug("对新浪微博返回<"+reqInfo+">的相关值格式化为struct：", resultStruct)

	return
}
