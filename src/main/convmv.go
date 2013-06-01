package main

import (
	"fmt"
	"strconv"
)

const (
	filePath    = "/media/FC3050D730509B0A/Movice"
	folder      = "FD69D5D671C9D5BA017AD395BE4A234B8CB8310E"
	fileName    = "西游降魔篇.BD1280高清中字.rmvb"
	fileNameOld = fileName + "_"
	fileSuffix  = ".\\!mv"
	feg         = 115
)

func main() {

	var fileList string

	for i := 0; i <= feg; i++ {
		fileList += fileNameOld + strconv.Itoa(i) + fileSuffix + " "
	}

	fmt.Println(cat "+fileList+" > "+fileName)
}
