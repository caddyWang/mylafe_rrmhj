package business

import (
	//"github.com/astaxie/beego"
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
			roleInfo.MtPicName = role[titles["mtpicname"]]

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

			sort, err2 := strconv.Atoi(role[titles["sort"]])
			if err2 != nil {
				sort = 999999
			}
			roleInfo.Sort = sort

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
			src.ProfileImg = d[titles["profileimg"]]
			src.Color = d[titles["color"]]
			src.ProfileText = d[titles["profiletext"]]
			src.MtPicName = d[titles["mtpicname"]]

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

			sort, err2 := strconv.Atoi(d[titles["sort"]])
			if err2 != nil {
				sort = 999999
			}
			src.Sort = sort

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
			src.ProfileImg = d[titles["profileimg"]]
			src.MtPicName = d[titles["mtpicname"]]

			sysrole, err1 := strconv.Atoi(d[titles["systemrole"]])
			if err1 != nil {
				sysrole = 0
			}
			src.SystemRole = sysrole

			dao.InsertScene(&src)
		}
	}
}

func InitRoleFaceInfo(roleData [][]string, titles map[string]int) {
	var roleInfo resource.SrcRoleFaceInfo

	for _, role := range roleData {
		if role[titles["optcode"]] == "1" {
			roleInfo.FaceName = role[titles["facename"]]
			roleInfo.PicName = role[titles["picname"]]
			roleInfo.ItemPicName = role[titles["itempicname"]]
			roleInfo.ProfileName = role[titles["profilename"]]
			roleInfo.ProfileImg = role[titles["profileimg"]]
			roleInfo.RoleName = role[titles["rolename"]]
			roleInfo.MtPicName = role[titles["mtpicname"]]

			sysrole, err1 := strconv.Atoi(role[titles["systemrole"]])
			if err1 != nil {
				sysrole = 0
			}
			roleInfo.SystemRole = sysrole

			sort, err2 := strconv.Atoi(role[titles["sort"]])
			if err2 != nil {
				sort = 999999
			}
			roleInfo.Sort = sort

			dao.InsertRoleFace(&roleInfo)
		}
	}
}

func InitRoleActionClothingInfo(roleData [][]string, titles map[string]int) {

	for _, role := range roleData {

		if role[titles["optcode"]] == "1" {
			var roleInfo resource.SrcRoleActionInfo
			var cl resource.SrcRoleClothingInfo

			roleInfo.ItemPicName = role[titles["itemactionpicname"]]
			roleInfo.ProfileName = role[titles["actionprofilename"]]
			roleInfo.ProfileImg = role[titles["actionprofileimg"]]
			roleInfo.ActionGroup = role[titles["actiongroup"]]
			roleInfo.RoleName = role[titles["rolename"]]
			roleInfo.MtPicName = role[titles["mtactionpicname"]]

			var clothing struct {
				ClothingName  string
				PicName       string
				ItemPicName   string
				ClothingGroup string
			}
			clothing.ClothingName = role[titles["clothingname"]]
			clothing.PicName = role[titles["picname"]]
			clothing.ItemPicName = role[titles["itemclothingpicname"]]
			clothing.ClothingGroup = role[titles["clothinggroup"]]

			roleInfo.Clothing = append(roleInfo.Clothing, clothing)

			sysrole, err1 := strconv.Atoi(role[titles["systemrole"]])
			if err1 != nil {
				sysrole = 0
			}
			roleInfo.SystemRole = sysrole

			sort, err2 := strconv.Atoi(role[titles["sort"]])
			if err2 != nil {
				sort = 999999
			}
			cl.Sort = sort

			cl.ClothingName = clothing.ClothingName
			cl.ItemPicName = clothing.ItemPicName
			cl.PicName = clothing.PicName
			cl.MtPicName = role[titles["mtclothingpicname"]]
			cl.ProfileName = role[titles["clothingprofilename"]]
			cl.ProfileImg = role[titles["clothingprofileimg"]]
			cl.ClothingGroup = clothing.ClothingGroup
			cl.RoleName = roleInfo.RoleName
			cl.SystemRole = sysrole

			dao.InsertRoleAction(&roleInfo)
			dao.InsertRoleClothing(&cl)
		}
	}
}
