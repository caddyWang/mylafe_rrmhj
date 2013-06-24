package business

import (
	"rrmhjbg.com/dao"
	"rrmhjbg.com/models"
)

func InitUserInfoBySinaWeibo(socialUser models.SocialUserInfo, platformName string) (uid string) {
	user := models.UserInfo{}

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
