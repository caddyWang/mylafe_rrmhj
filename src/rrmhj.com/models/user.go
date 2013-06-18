package models

import (
	//"labix.org/v2/mgo/bson"
	"time"
)

type UserBase struct {
	Id         string "_id"
	UserName   string
	ProfileImg string
}

type UserInfo struct {
	Id         string "_id"
	UserName   string
	ProfileImg string

	ProfileLargeImg string
	Gender          string
	Province        string
	City            string
	Location        string
	CreateTime      time.Time

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
