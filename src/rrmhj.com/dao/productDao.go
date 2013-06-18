package dao

/************************************************************************************
//
// Desc		:	作品数据Dao
// Records	:	2013-06-18	Wangdj	新建文件；增加函数"InitUserInfoBySinaWeibo"
//
************************************************************************************/
import (
	"github.com/astaxie/beego"
	"labix.org/v2/mgo/bson"
	"rrmhj.com/conf"
	"rrmhj.com/models"
)

const proInfo = "productInfo"

func GetProductListByPage(pageIndex int) (proList []models.Product, count int) {
	if pageIndex < 0 {
		pageIndex = 0
	}

	pageSize := conf.PageSize
	proList = []models.Product{}

	count, err := FindList(bson.M{"iflag": 0}, &proList, proInfo, pageIndex, pageSize, "-posttime")
	if err != nil {
		beego.Error("查询漫画列表数据出错：", err)
	}

	return proList, count
}
