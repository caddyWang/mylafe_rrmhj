package conf

/************************************************************************************
//
// Desc		:	从配置文件中读取配置信息
//
************************************************************************************/

import (
	"github.com/astaxie/beego"
)

//mongoDB
var ConnAddr = beego.AppConfig.String("conn_addr")
var DefDBName = beego.AppConfig.String("db_name")
var ResourceDBName = beego.AppConfig.String("db_name_resource") //资源数据库

var PageSize, _ = beego.AppConfig.Int("page_size")
var ImgUrl = beego.AppConfig.String("img_url")
