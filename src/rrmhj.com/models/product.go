package models

import (
	"time"
)

type Product struct {
	Pid      int
	ImgPath  string
	Author   UserBase
	PostTime time.Time
	Desc     string
	UpNum    int
	DownNum  int
	Comments []Comment
}

type Comment struct {
	Cid         int
	Reviewer    UserBase
	PostTime    time.Time
	CommentDesc string
}
