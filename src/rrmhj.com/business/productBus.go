package business

/************************************************************************************
//
// Desc		:	与漫画作品有关的业务功能
// Records	:	2013-06-07	Wangdj	新建文件；增加函数"QueryProductsList"
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
