package business

/************************************************************************************
//
// Desc		:	与会员相关的业务功能
// Records	:	2013-06-14	Wangdj	新建文件；增加函数"CheckLogin"
//
************************************************************************************/

type GetSession func(key interface{}) interface{}

//验证用户是否登录
func CheckLogin(gs GetSession) bool {
	uid := gs("uid")

	if uid == nil {
		return false
	}

	if uid == "" {
		return false
	}

	return true
}
