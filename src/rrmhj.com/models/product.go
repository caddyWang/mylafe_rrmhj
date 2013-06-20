package models

import (
	"time"
)

type Product struct {
	Pid        string "_id"
	ImgPath    string
	Author     UserBase
	PostTime   time.Time
	Desc       string
	UpNum      int
	DownNum    int
	CommentNum int
	Iflag      int //0为审核通过可查看，-1为不可查看
}

type ProductUseHtml struct {
	Product
	UpNumScript   string
	DownNumScript string
}

type Comment struct {
	Cid         string "_id"
	Proid       string
	Reviewer    UserBase
	PostTime    time.Time
	CommentDesc string
}
