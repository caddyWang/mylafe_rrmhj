package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	. "github.com/qiniu/api/conf"
	"github.com/qiniu/api/io"
	"github.com/qiniu/api/rs"
	"github.com/qiniu/rpc"
	"rrmhj.com/business"
	"rrmhj.com/models"
	"strconv"
	"strings"
	"time"
)

var logger rpc.Logger
var ret io.PutRet
var extra = &io.PutExtra{}

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

type PutImgCloudController struct {
	beego.Controller
}

func (this *PutImgCloudController) Post() {
	r := this.Ctx.Request
	r.ParseMultipartForm(32 << 20)
	file, hander, err := r.FormFile("qqfile")
	if err != nil {
		beego.Error("...", err)
		return
	}
	defer file.Close()

	ACCESS_KEY = "hQDkHrrIArdxEcthaHNSa4GqjKm0LD1G5HbcTbH_"
	SECRET_KEY = "X8JWFgipMFfqifIoA7KTZUcHS4iMq4I9N9tfR1N3"

	uptoken := uptoken("rrmhj")
	key := strconv.FormatInt(time.Now().UnixNano(), 10)
	suffix := hander.Filename[strings.LastIndex(hander.Filename, "."):]

	err = io.Put(logger, &ret, uptoken, key+suffix, file, extra)
	if err != nil {
		beego.Error("...", err)
		return
	}

	var rtn struct {
		io.PutRet
		Success bool `json:"success"`
	}
	rtn.Hash, rtn.Key, rtn.Success = ret.Hash, ret.Key, true
	infoJson, err1 := json.Marshal(rtn)
	if err1 != nil {
		beego.Error("数据格式化成JSON出错！", err1)
	}
	beego.Debug(rtn)
	beego.Debug(string(infoJson))

	this.Ctx.WriteString(string(infoJson))
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
