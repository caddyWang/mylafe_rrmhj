package business

import (
	"rrmhjbg.com/dao"
	"rrmhjbg.com/models"
)

func InitUserInfoBySinaWeibo(socialUser models.SocialUserInfo, platformName, rrmhjUid string) (uid string) {
	user := models.UserInfo{Id: rrmhjUid}

	switch platformName {
	case models.SinaWeibo:
		user.SinaWeibo = []models.SocialUserInfo{socialUser}

	case models.TencWeibo:
		user.TencWeibo = []models.SocialUserInfo{socialUser}

	case models.QQZone:
		user.QQZone = []models.SocialUserInfo{socialUser}

	case models.RenRenSNS:
		user.RenRenSNS = []models.SocialUserInfo{socialUser}
	}

	return dao.InitUserInfoBySinaWeibo(&user, platformName)
}

func GetAllUsers(pageIndex, pageSize int, user *[]models.UserInfo) (icount, optCode int) {
	count, err := dao.GetAllUsers(pageIndex, pageSize, "", user)
	if err != nil {
		optCode = -1
	}

	return count, optCode
}

func UpdateUser(user *models.UserInfo) {
	dao.UpateUser(user)
}
