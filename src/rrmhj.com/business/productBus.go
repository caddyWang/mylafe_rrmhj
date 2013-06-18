package business

/************************************************************************************
//
// Desc		:	与漫画作品有关的业务功能
// Records	:	2013-06-07	Wangdj	新建文件；增加函数"QueryProductsList"
//
************************************************************************************/

import (
	"rrmhj.com/models"
	"time"
)

//根据分页参数获取漫画作品列表（Wangdj 2013-06-07	）
func QueryProductsList(pageIndex, pageSize int) (prolist []models.Product, err error) {

	if pageIndex--; pageIndex < 0 {
		pageIndex = 0
	}

	prolist = []models.Product{}

	for i := 0; i < pageSize; i++ {
		p := models.Product{i, "test/1369388229312.jpg", models.UserBase{"1", "dadairen", ""}, time.Now(), "麻麻最近迷上了画漫画，她把家庭生活中的故事通过卡通人物讲了出来，50岁的年龄却有18岁的style，赞！", 155, -23, nil}
		comments := []models.Comment{{i, models.UserBase{"1", "haha", "test/user1.jpg"}, time.Now(), "话说，貌似不应该用炒菜勺拖地吧太可爱了!"}, {2, models.UserBase{"2", "麦麦粉", ""}, time.Now(), "太可爱了!"}}
		p.Comments = comments

		prolist = append(prolist, p)
	}

	return prolist, nil

}
