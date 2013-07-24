package tools

import (
	"archive/zip"
	"bytes"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"regexp"
)

//过滤掉URL中的一些非法字符
func FilterURL(origin string) (dest string) {

	re, _ := regexp.Compile("[a-zA-Z0-9/-/_/:/.//]*")
	one := re.Find([]byte(origin))

	return (string(one))

}

//2013/07/23 Wangdj 根据用户请求的信息，生成可下载的ZIP文件包
func GencZip(srcFiles []string, url string, confFileContent []byte) (zipFile []byte) {

	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)

	for _, file := range srcFiles {
		f, err := w.Create(file)
		handlerErr("GencZip", err)

		resp, err1 := http.Get(url + file)
		handlerErr("GencZip", err1)

		result, err2 := ioutil.ReadAll(resp.Body) //取出主体的内容
		defer resp.Body.Close()
		handlerErr("GencZip", err2)
		beego.Debug("resp.Body", len(result))

		_, err = f.Write(result)
		handlerErr("GencZip", err)
	}

	//add config file
	f, err := w.Create("imgprofile")
	handlerErr("GencZip", err)
	_, err = f.Write(confFileContent)
	handlerErr("GencZip", err)

	//close zip
	err = w.Close()
	handlerErr("GencZip", err)
	return buf.Bytes()
}

func handlerErr(funcName string, err error) {
	if err != nil {
		beego.Error(funcName, err)
	}
}