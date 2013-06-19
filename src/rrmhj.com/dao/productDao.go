package dao

/************************************************************************************
//
// Desc		:	作品数据Dao
// Records	:	2013-06-18	Wangdj	新建文件；增加函数"InitUserInfoBySinaWeibo"
//				2013-06-19	Wangdj	增加函数"SaveProComment"
//
************************************************************************************/
import (
	"github.com/astaxie/beego"
	"labix.org/v2/mgo/bson"
	"rrmhj.com/conf"
	"rrmhj.com/models"
	"time"
)

const proInfo = "productInfo"
const commInfo = "commentInfo"

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

//读取某个漫画下的所有评论(Wangdj 2013-06-19)
func GetProComment(pid string) (commentList []models.Comment, count int, err error) {
	commentList = []models.Comment{}

	count, err = FindList(bson.M{"proid": pid}, &commentList, commInfo, 0, 100, "-posttime")
	if err != nil {
		beego.Error("读取评论数时出错： proid=", pid, err)
	}
	return
}

//保存评论
func SaveProComment(comment *models.Comment) (err error) {
	comment.Cid = bson.NewObjectId().Hex()
	comment.PostTime = time.Now()

	err = Insert(commInfo, &comment)
	if err != nil {
		beego.Error("保存评论出错：", *comment, err)
	}

	err = Update(proInfo, bson.M{"_id": comment.Proid}, bson.M{"$inc": bson.M{"commentnum": 1}})
	if err != nil {
		beego.Error("增加评论数时出错：", *comment, err)
	}

	return
}
