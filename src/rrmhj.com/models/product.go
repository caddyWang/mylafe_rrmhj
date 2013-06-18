package models

import (
	"time"
)

type Product struct {
	Pid      string "_id"
	ImgPath  string
	Author   UserBase
	PostTime time.Time
	Desc     string
	UpNum    int
	DownNum  int
	Comments []Comment
	Iflag    int //0为审核通过可查看，-1为不可查看
}

type Comment struct {
	Cid         int
	Reviewer    UserBase
	PostTime    time.Time
	CommentDesc string
}
