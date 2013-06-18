package models

import (
	"labix.org/v2/mgo/bson"
)

type UserBase struct {
	Id         bson.ObjectId "_id"
	UserName   string
	ProfileImg string
}

type UserInfo struct {
	Id         bson.ObjectId "_id"
	UserName   string
	ProfileImg string

	Gender   string
	Province string
	City     string
	Location string

	SinaWeibo SinaWeiboUserInfo
}

type SinaWeiboUserInfo struct {
	SnUid         int64
	SnUserName    string
	SnProfileImg  string
	SnAvaterLarge string
	SnProfileUrl  string
	SnDescription string
}
