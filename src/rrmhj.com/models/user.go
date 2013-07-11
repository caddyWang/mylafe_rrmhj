package models

import (
	"time"
)

const (
	SinaWeibo = "sina"
	TencWeibo = "tencent"
	QQZone    = "qqzone"
	RenRenSNS = "renren"
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
	LikePro         []string

	SinaWeibo []SocialUserInfo
	TencWeibo []SocialUserInfo
	QQZone    []SocialUserInfo
	RenRenSNS []SocialUserInfo
}

type SocialUserInfo struct {
	Uid        string
	UserName   string
	ProfileImg string
	ProfileUrl string

	Gender      string
	Province    string
	City        string
	Location    string
	AvaterLarge string
	Description string
}
