package dao

/************************************************************************************
//
// Desc		:	数据库连接函数
// Records	:	2013-06-18	Wangdj	新建文件；增加函数"GetDBDef"、"getDBConn"
//
************************************************************************************/

import (
	"errors"
	"github.com/astaxie/beego"
	"labix.org/v2/mgo"
	"rrmhj.com/conf"
)

const (
	findOneOpt = iota
	insertOneOpt
)

var dbNameDef = conf.DefDBName

func FindOne(query interface{}, result interface{}, collectionName string) (err error) {
	return optDB(findOneOpt, dbNameDef, collectionName, query, result)
}

func InsertOne(collectionName string, doc interface{}) (err error) {
	return optDB(insertOneOpt, dbNameDef, collectionName, nil, doc)
}

//获取数据库
func optDB(optType int, dbname string, collectionName string, query interface{}, docs ...interface{}) (err error) {

	session, err := getDBConnSession()
	if err != nil {
		beego.Critical("数据库连接出错：", err)
		panic(err)
	}
	defer session.Close()

	if len(docs) <= 0 {
		beego.Error("docs必须大于零")
		return errors.New("docs必须大于零")
	}

	c := session.DB(dbname).C(collectionName)
	switch optType {
	case findOneOpt:
		err = c.Find(query).One(docs[0])
	case insertOneOpt:
		err = c.Insert(docs[0])
	}

	return
}

//获取数据库连接
func getDBConnSession() (session *mgo.Session, err error) {

	session, err = mgo.Dial(conf.ConnAddr)
	return
}
