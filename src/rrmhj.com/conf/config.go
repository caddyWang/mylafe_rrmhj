package conf

/************************************************************************************
//
// Desc		:	从配置文件中读取配置信息
// Records	:	2013-06-18	Wangdj	新建文件；增加函数"GetDBDef"、"getDBConn"
//
************************************************************************************/

import (
	"github.com/astaxie/beego"
)

var StaticFileURL = beego.AppConfig.String("static_file_url") //js,css,img等静态文件的默认路径
var PageSize, _ = beego.AppConfig.Int("page_size")            //每页显示条数
var DefProfileImg = beego.AppConfig.String("def_profile_img") //用户没有设置头像时的，默认头像

//Sina weibo API's config params
var SinaRedirectUri = beego.AppConfig.String("sina_redirect_uri")                      //用新浪微博帐户登录后的，回调页面
var SinaOauth2AccesstokenAddr = beego.AppConfig.String("sina_oauth2_accesstoken_addr") //请求access token的地址
var SinaClientId = beego.AppConfig.String("sina_client_id")                            //新浪appKey
var SinaClientSecret = beego.AppConfig.String("sina_client_secret")                    //新浪appSecret
var SinaGrantType = beego.AppConfig.String("sina_grant_type")                          //请求access token时grant类型
var SinaUserShowAddr = beego.AppConfig.String("sina_user_show_addr")                   //请求用户信息的地址

//mongoDB
var ConnAddr = beego.AppConfig.String("conn_addr")
var DefDBName = beego.AppConfig.String("db_name")
