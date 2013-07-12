package business

/************************************************************************************
//
// Desc		:	与漫画作品有关的业务功能
// Records	:	2013-06-07	Wangdj	新建文件；增加函数"QueryProductsList"
//				2013-06-19	Wangdj	增加函数"SaveProductComment"
//
************************************************************************************/

import (
	"github.com/astaxie/beego"
	"net/http"
	"rrmhj.com/dao"
	"rrmhj.com/models"
)

//根据分页参数获取漫画作品列表（Wangdj 2013-06-07	）
func QueryProductsList(pageIndex int, req *http.Request) (proHtmllist []models.ProductUseHtml, count int) {

	if pageIndex < 0 {
		pageIndex = 0
	}

	prolist, count := dao.GetProductListByPage(pageIndex)
	return setUpNumScript(prolist, req), count
}

//读取某个漫画下的所有评论(Wangdj 2013-06-19)
func GetProComment(pid string) (commentList []models.Comment, count int, err error) {
	var list []models.Comment
	commentList = []models.Comment{}
	list, count, err = dao.GetProComment(pid)

	for _, comment := range list {
		comment.CommentDesc = beego.Html2str(comment.CommentDesc)
		comment.Reviewer.ProfileImg = DefaultHeadImg(comment.Reviewer.ProfileImg)
		commentList = append(commentList, comment)
	}

	return commentList, count, err
}

//保存漫画评论（Wangdj 2013-06-19）
func SaveProductComment(comment *models.Comment, gs GetSession) (err error) {
	comment.Reviewer = GetSessinUserBase(gs)
	err = dao.SaveProComment(comment)

	comment.Reviewer.ProfileImg = DefaultHeadImg(comment.Reviewer.ProfileImg)

	return
}

//2013/06/20 Wangdj 新建：更新用户踩或顶
//2013/07/10 Wangdj 修改：只保留“顶”功能，并增加顶的表情选择
func UpdateProUporDown(proId, dingface string) {
	dao.UpdateProUporDown(proId, dingface)
}

//2013/07/11 Wangdj 新建：获取指定用户发布的作品集
//2013/07/12 Wangdj 修改：根据当前Cookies判断当前用户是否已经顶过各个作品，并添加显示样式
func GetProductsByUid(ctx *beego.Controller, pageIndex int) (proHtmllist []models.ProductUseHtml, count int) {
	var proList []models.Product

	if pageIndex < 0 {
		pageIndex = 0
	}

	uid := ctx.GetSession("uid")
	if uid == nil || uid == "" {
		return []models.ProductUseHtml{}, 0
	}

	proList, count = dao.GetProductsByUid(uid.(string), pageIndex)
	return setUpNumScript(proList, ctx.Ctx.Request), count
}

func GetLikeProByUid(ctx *beego.Controller, pageIndex int) (proHtmllist []models.ProductUseHtml, count int) {
	var proList []models.Product

	if pageIndex < 0 {
		pageIndex = 0
	}

	uid := ctx.GetSession("uid")
	if uid == nil || uid == "" {
		return []models.ProductUseHtml{}, 0
	}

	pidArr := dao.GetUserLikeProduct(uid.(string))
	proList, count = dao.GetProductLikeByPidArr(pidArr, pageIndex)
	return setUpNumScript(proList, ctx.Ctx.Request), count
}

//2013/07/12 Wangdj 新建：删除用户指定的作品
func DelProductByUid(uid, pid string) (err error) {
	return dao.DelProductByUid(uid, pid)
}

//2013/07/12 Wangdj 新增：内部公共函数，根据当前Cookies判断当前用户是否已经顶过各个作品，并添加显示样式
func setUpNumScript(proList []models.Product, req *http.Request) (proHtmllist []models.ProductUseHtml) {
	proHtmllist = []models.ProductUseHtml{}

	for _, pro := range proList {
		proHtml := models.ProductUseHtml{}
		proHtml.Pid, proHtml.ImgPath, proHtml.Author, proHtml.PostTime, proHtml.Desc, proHtml.UpNum, proHtml.DownNum, proHtml.CommentNum = pro.Pid, pro.ImgPath, pro.Author, pro.PostTime, pro.Desc, pro.UpNum, pro.DownNum, pro.CommentNum
		proHtml.UpNumScript, proHtml.DownNumScript = "up", "down"
		cookie, err := req.Cookie(pro.Pid)
		if err == nil && cookie != nil {
			val := cookie.Value
			if val == "1" {
				proHtml.UpNumScript = "ding_disabled"
			}
		}

		proHtmllist = append(proHtmllist, proHtml)
	}

	return
}
