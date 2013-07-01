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
