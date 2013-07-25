package business

import (
	"rrmhjbg.com/dao"
	"rrmhjbg.com/models/resource"
	"strconv"
)

func InitRoleInfo(roleData [][]string, titles map[string]int) {
	var roleInfo resource.SrcRoleInfo

	for _, role := range roleData {
		if role[titles["optcode"]] == "1" {
			roleInfo.RoleName = role[titles["rolename"]]
			roleInfo.PicName = role[titles["picname"]]
			roleInfo.ItemPicName = role[titles["itempicname"]]
			roleInfo.DefaultFace = role[titles["defaultface"]]
			roleInfo.DefaultClothing = role[titles["defaultclothing"]]
			roleInfo.ProfileName = role[titles["profilename"]]
			roleInfo.ProfileImg = role[titles["profileimg"]]
			roleInfo.UndownImg = role[titles["undownimg"]]
			roleInfo.ProfileText = role[titles["profiletext"]]

			dir, err := strconv.Atoi(role[titles["direction"]])
			if err != nil {
				dir = 1
			}
			roleInfo.Direction = dir

			sysrole, err1 := strconv.Atoi(role[titles["systemrole"]])
			if err1 != nil {
				sysrole = 0
			}
			roleInfo.SystemRole = sysrole

			dao.InsertRole(&roleInfo)
		}
	}
}

func InitDialogInfo(data [][]string, titles map[string]int) {
	var src resource.SrcDialogInfo

	for _, d := range data {
		if d[titles["optcode"]] == "1" {
			src.DialogName = d[titles["dialogname"]]
			src.PicName = d[titles["picname"]]
			src.ItemPicName = d[titles["itempicname"]]
			src.ProfileName = d[titles["profilename"]]
			src.Color = d[titles["color"]]
			src.ProfileText = d[titles["profiletext"]]

			dir, err := strconv.Atoi(d[titles["direction"]])
			if err != nil {
				dir = 1
			}
			src.Direction = dir

			sysrole, err1 := strconv.Atoi(d[titles["systemrole"]])
			if err1 != nil {
				sysrole = 0
			}
			src.SystemRole = sysrole

			dao.InsertDialog(&src)
		}
	}
}

func InitSceneInfo(data [][]string, titles map[string]int) {
	var src resource.SrcSceneInfo

	for _, d := range data {
		if d[titles["optcode"]] == "1" {
			src.SceneName = d[titles["scenename"]]
			src.PicName = d[titles["picname"]]
			src.ItemPicName = d[titles["itempicname"]]
			src.ProfileName = d[titles["profilename"]]
			src.ProfileText = d[titles["profiletext"]]

			sysrole, err1 := strconv.Atoi(d[titles["systemrole"]])
			if err1 != nil {
				sysrole = 0
			}
			src.SystemRole = sysrole

			dao.InsertScene(&src)
		}
	}
}
