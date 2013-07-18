package business

import (
	"fmt"
	"rrmhjbg.com/conf"
	"rrmhjbg.com/dao"
	"rrmhjbg.com/models/jsonmodels"
	"rrmhjbg.com/models/resource"
	"strconv"
)

var (
	roleInfoList   []resource.SrcRoleInfo
	dialogInfoList []resource.SrcDialogInfo
	sceneInfoList  []resource.SrcSceneInfo
)

//2013/07/18 Wangdj 新增：获取当前分页下的资源信息，验证数据合法性，并转换成json
func ShowSrcInfoByPage(pageIndex, pageSize, srcType, uid string, showResList *jsonmodels.ShowResList) {
	var count int
	var err, err1 error
	var downRoleInfo string

	index, size := filterPageInfo(pageIndex, pageSize)
	showResList.PageIndex = strconv.Itoa(index)
	showResList.PageSize = strconv.Itoa(size)
	showResList.ImgSuffix = conf.ImgUrl

	userDownd := resource.SrcUserDownloaded{}
	err1 = dao.GetDownloadInfoByUid(uid, &userDownd)

	switch srcType {
	case "1":
		roleInfoList = []resource.SrcRoleInfo{}
		count, err = dao.ShowRoleInfoByPage(index, size, &roleInfoList)

	case "2":
		dialogInfoList = []resource.SrcDialogInfo{}
		count, err = dao.ShowDialogInfoByPage(index, size, &dialogInfoList)

	case "3":
		sceneInfoList = []resource.SrcSceneInfo{}
		count, err = dao.ShowSceneInfoByPage(index, size, &sceneInfoList)
	}

	if err != nil || err1 != nil {
		showResList.OptCode = "-1"
	} else {
		showResList.OptCode = "0"
		showResList.ListCount = strconv.Itoa(count)

		switch srcType {
		case "1":
			downRoleInfo = fmt.Sprint(userDownd.RoleInfo)
			for _, rec := range roleInfoList {
				showResList.ListArry = append(showResList.ListArry, rec.GetRes(downRoleInfo))
			}

		case "2":
			downRoleInfo = fmt.Sprint(userDownd.DialogInfo)
			for _, rec := range dialogInfoList {
				showResList.ListArry = append(showResList.ListArry, rec.GetRes(downRoleInfo))
			}

		case "3":
			downRoleInfo = fmt.Sprint(userDownd.SceneInfo)
			for _, rec := range sceneInfoList {
				showResList.ListArry = append(showResList.ListArry, rec.GetRes(downRoleInfo))
			}
		}

	}
}

//2013/07/18 Wangdj 对分页条件进行过滤
func filterPageInfo(pageIndex, pageSize string) (index, size int) {
	var err, err1 error
	index, err = strconv.Atoi(pageIndex)
	size, err1 = strconv.Atoi(pageSize)

	if err != nil || index < 1 {
		index = 1
	}

	if err1 != nil || size < 1 {
		size = conf.PageSize
	}

	return
}
