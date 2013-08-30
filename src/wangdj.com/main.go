package main

import (
	//"bytes"
	"fmt"
	//"os"
	//"os/exec"
	"strconv"
	"time"
)

const (
	origin = 0
)

// cat 疯狂原始人_2013_720p.mkv_0.!mv 疯狂原始人_2013_720p.mkv_1.!mv > 疯狂原始人_2013_720p.mkv
func main() {

	/*
	args := os.Args
	if len(args) != 3 {
		fmt.Println("args not good!")
		return
	}

	fileNum, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("args not good!")
		return
	}

	fileName := args[1]
	cmd := []string{"cat"}

	for i := origin; i <= fileNum; i++ {
		file := fileName + "_" + strconv.Itoa(i) + ".\\!mv"
		cmd = append(cmd, file)
	}
	cmd = append(cmd, ">")
	cmd = append(cmd, fileName)

	fmt.Println(cmd)

	
		cat := exec.Command("cat", cmd...)
		var out bytes.Buffer
		cat.Stdout = &out
		err = cat.Run()
		if err != nil {
			fmt.Println("command err : ", err.Error())
		}

	*/

	for i:=0; i<=10000; i++ {
		fmt.Println(getUniqID())
	}

}

func getUniqID() (string){
	uid := strconv.FormatInt(time.Now().UnixNano(), 10)
    return uid
}
