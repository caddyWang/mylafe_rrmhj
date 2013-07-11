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
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"rrmhj.com/conf"
	"rrmhj.com/models"
	"time"
)

const proInfo = "productInfo"
const commInfo = "commentInfo"

func GetProductListByPage(pageIndex int) (proList []models.Product, count int) {

	pageSize := conf.PageSize
	proList = []models.Product{}

	count, err := FindList(bson.M{"iflag": 0}, &proList, proInfo, pageIndex*pageSize, pageSize, "-posttime")
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

//2013/06/20 Wangdj 新建：更新用户踩或顶
//2013/07/10 Wangdj 修改：只保留“顶”功能，并增加顶的表情选择
func UpdateProUporDown(proId, dingface string) {
	var change interface{}

	change = bson.M{"$inc": bson.M{"upnum": 1, "dingface." + dingface: 1}}

	beego.Debug("proId=", proId)
	beego.Debug("dingface=", dingface)
	beego.Debug("change=", change)

	err := Update(proInfo, bson.M{"_id": proId}, change)
	if err != nil {
		beego.Error("更新用户顶时出错：proId=", proId, err)
	}
}

//2013/07/11 Wangdj 新建：获取指定用户发布的作品集
func GetProductsByUid(uid string, pageIndex int) (proList []models.Product, count int) {

	pageSize := 10
	proList = []models.Product{}

	count, err := FindList(bson.M{"author.id": uid}, &proList, proInfo, pageIndex*pageSize, pageSize, "-posttime")
	if err == mgo.ErrNotFound {
		return []models.Product{}, 0
	} else if err != nil {
		beego.Error("查询指定用户漫画列表数据出错：Uid=", uid, err)
		return []models.Product{}, 0
	}

	return proList, count
}
