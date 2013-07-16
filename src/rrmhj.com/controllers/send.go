package controllers

import (
	"github.com/astaxie/beego"
	. "github.com/qiniu/api/conf"
	"github.com/qiniu/api/rs"
	"rrmhj.com/business"
	"rrmhj.com/models"
	"time"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplNames = "send/login.tpl"
}

func (this *LoginController) Post() {
	userName, password := this.GetString("user"), this.GetString("password")

	beego.Debug(userName, password)
	beego.Debug(beego.AppConfig.String("username"), beego.AppConfig.String("password"))
	if userName != beego.AppConfig.String("username") || password != beego.AppConfig.String("password") {
		this.Redirect("/send/login?code=1", 302)
		return
	}

	this.SetSession("userName", userName)

	this.Redirect("/send/pro", 302)
}

type UploadController struct {
	beego.Controller
}

func (this *UploadController) Get() {

	if this.GetSession("userName") == nil {
		this.Redirect("/send/login", 302)
		return
	}

	ACCESS_KEY = "hQDkHrrIArdxEcthaHNSa4GqjKm0LD1G5HbcTbH_"
	SECRET_KEY = "X8JWFgipMFfqifIoA7KTZUcHS4iMq4I9N9tfR1N3"

	this.Data["Uptoken"] = uptoken("rrmhj")
	this.Data["Key"] = time.Now().UnixNano()

	user := this.GetSession("userName")
	if user.(string) == "" {
		this.Redirect("/send/login", 302)
		return
	}

	this.Data["IsLogin"] = business.CheckLogin(this.GetSession)
	business.LoginedUserInfo(&this.Controller)

	this.TplNames = "send/upload.tpl"
}

func (this *UploadController) Post() {
	imgName, descript := this.GetString("imgName"), this.GetString("descript")

	if this.GetSession("uid") != nil {
		uname := this.GetSession("uname").(string)
		uid := this.GetSession("uid").(string)
		userImg := this.GetSession("uprofileimg").(string)

		if uname == "" || uid == "" {
			return
		} else {
			product := models.Product{ImgPath: imgName, Desc: descript}
			product.Author.Id = uid
			product.Author.UserName = uname
			product.Author.ProfileImg = userImg

			business.SaveProduct(&product)
		}
	} else {
		return
	}

	defer this.Redirect("/send/pro", 302)
}

func uptoken(bucketName string) string {

	putPolicy := rs.PutPolicy{
		Scope: bucketName,
		//CallbackUrl: callbackUrl,
		//CallbackBody:callbackBody,
		//ReturnUrl:   returnUrl,
		//ReturnBody:  returnBody,
		//AsyncOps:    asyncOps,
		//EndUser:     endUser,
		//Expires:     expires,
	}
	return putPolicy.Token(nil)
}
