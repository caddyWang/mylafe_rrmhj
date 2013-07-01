package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	filePath = "/media/FC3050D730509B0A/Movice"
	folder   = "FD69D5D671C9D5BA017AD395BE4A234B8CB8310E"
	//fileName    = "西游降魔篇.BD1280高清中字.rmvb"
	//fileNameOld = fileName + "_"
	fileSuffix = ".\\!mv"
	//feg         = 115
)

func main() {

	var fileList, fileName, fileNameOld string

	args := os.Args

	if len(args) != 2 {
		fmt.Println("Error : 参数不正确！正确格式如：convmv 西游降魔篇.BD1280高清中字.rmvb 115")
		return
	}

	fileName = args[0]
	fileNameOld = fileName + "_"
	feg, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Error : 参数2：文件数必须为整数！")
		return
	}

	for i := 0; i <= feg; i++ {
		fileList += fileNameOld + strconv.Itoa(i) + fileSuffix + " "
	}

	fmt.Println("cat " + fileList + " > " + fileName)
}
