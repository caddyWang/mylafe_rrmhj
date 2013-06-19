package business

/************************************************************************************
//
// Desc		:	与漫画作品有关的业务功能
// Records	:	2013-06-07	Wangdj	新建文件；增加函数"QueryProductsList"
//				2013-06-19	Wangdj	增加函数"SaveProductComment"
//
************************************************************************************/

import (
	"rrmhj.com/dao"
	"rrmhj.com/models"
)

//根据分页参数获取漫画作品列表（Wangdj 2013-06-07	）
func QueryProductsList(pageIndex int) (prolist []models.Product) {
	prolist, _ = dao.GetProductListByPage(pageIndex)
	return
}

//读取某个漫画下的所有评论(Wangdj 2013-06-19)
func GetProComment(pid string) (commentList []models.Comment, count int, err error) {
	return dao.GetProComment(pid)
}

//保存漫画评论（Wangdj 2013-06-19）
func SaveProductComment(comment *models.Comment, gs GetSession) (err error) {
	comment.Reviewer = GetSessinUserBase(gs)
	err = dao.SaveProComment(comment)
	comment.Reviewer.ProfileImg = DefaultHeadImg(comment.Reviewer.ProfileImg)

	return
}
