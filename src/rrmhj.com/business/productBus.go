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
	prolist, count := dao.GetProductListByPage(pageIndex)
	proHtmllist = []models.ProductUseHtml{}

	for _, pro := range prolist {
		proHtml := models.ProductUseHtml{}
		proHtml.Pid, proHtml.ImgPath, proHtml.Author, proHtml.PostTime, proHtml.Desc, proHtml.UpNum, proHtml.DownNum, proHtml.CommentNum = pro.Pid, pro.ImgPath, pro.Author, pro.PostTime, pro.Desc, pro.UpNum, pro.DownNum, pro.CommentNum
		proHtml.UpNumScript, proHtml.DownNumScript = "up", "down"
		cookie, err := req.Cookie(pro.Pid)
		if err == nil && cookie != nil {
			val := cookie.Value
			if val == "1" {
				proHtml.UpNumScript = "upselected btn-warning disabled"
				proHtml.DownNumScript = "downselected disabled"
			} else if val == "-1" {
				proHtml.UpNumScript = "upselected disabled"
				proHtml.DownNumScript = "downselected btn-warning disabled"
			}
		}

		proHtmllist = append(proHtmllist, proHtml)
	}

	return
}

//读取某个漫画下的所有评论(Wangdj 2013-06-19)
func GetProComment(pid string) (commentList []models.Comment, count int, err error) {
	commentList, count, err = dao.GetProComment(pid)
	for _, comment := range commentList {
		comment.CommentDesc = beego.Html2str(comment.CommentDesc)
	}
	return
}

//保存漫画评论（Wangdj 2013-06-19）
func SaveProductComment(comment *models.Comment, gs GetSession) (err error) {
	comment.Reviewer = GetSessinUserBase(gs)
	err = dao.SaveProComment(comment)
	comment.Reviewer.ProfileImg = DefaultHeadImg(comment.Reviewer.ProfileImg)

	return
}

//更新用户踩或顶(Wangdj 2013-06-20)
func UpdateProUporDown(proId string, optValue int) {
	dao.UpdateProUporDown(proId, optValue)
}
