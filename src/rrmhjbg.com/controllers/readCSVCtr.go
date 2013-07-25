package controllers

import (
	"encoding/csv"
	"github.com/astaxie/beego"
	"io"
	"os"
	"rrmhjbg.com/business"
)

type InitDataController struct {
	beego.Controller
}

func (this *InitDataController) Get() {
	readCSV("role", business.InitRoleInfo)
	readCSV("dialog", business.InitDialogInfo)
	readCSV("scene", business.InitSceneInfo)
}

func readCSV(fileName string, initFunc func([][]string, map[string]int)) {
	f, err := os.Open("./temp/" + fileName + ".csv")
	if err != nil {
		beego.Error(err)
		return
	}
	defer f.Close()

	var datas [][]string
	reader := csv.NewReader(f)
	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			beego.Error(err)
		}

		datas = append(datas, record)
	}

	if len(datas) > 0 {
		titles := make(map[string]int)
		for i, t := range datas[0] {
			titles[t] = i
		}

		initFunc(datas, titles)
	}

}
