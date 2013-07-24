package dao

/************************************************************************************
//
// Desc		:	数据库连接函数
// Records	:	2013-06-18	Wangdj	新建文件；增加函数"GetDBDef"、"getDBConn"
//
************************************************************************************/

import (
	"github.com/astaxie/beego"
	"labix.org/v2/mgo"
	"rrmhjbg.com/conf"
)

const (
	findOneOpt = iota
	findListOpt
)

var DBName = conf.ResourceDBName
var DBNameForUser = conf.DefDBName

func FindOne(query, result interface{}, collectionName string) (err error) {
	return FindOneDB(query, result, collectionName, DBName)
}
func FindOneDB(query, result interface{}, collectionName, dbName string) (err error) {
	_, err = queryOpt(findOneOpt, dbName, collectionName, query, result, "", 0, 0)
	return
}

func FindList(query, result interface{}, collectionName string, skip, limit int, sort string) (count int, err error) {
	return FindListDB(query, result, collectionName, DBName, skip, limit, sort)
}

func FindListDB(query, result interface{}, collectionName, dbName string, skip, limit int, sort string) (count int, err error) {
	if skip < 0 {
		skip = 0
	}

	return queryOpt(findListOpt, dbName, collectionName, query, result, sort, skip, limit)
}

func Insert(collectionName string, doc interface{}) (err error) {
	return InsertDB(collectionName, DBName, doc)
}

func InsertDB(collectionName, dbName string, doc interface{}) (err error) {

	session, err := getDBConnSession()
	if err != nil {
		beego.Critical("数据库连接出错：", err)
		panic(err)
	}
	defer session.Close()

	c := session.DB(dbName).C(collectionName)
	err = c.Insert(doc)

	return
}

func Update(collectionName string, selector, change interface{}) (err error) {
	return UpdateDB(collectionName, DBName, selector, change)
}

func UpdateDB(collectionName, dbName string, selector, change interface{}) (err error) {

	session, err := getDBConnSession()
	if err != nil {
		beego.Critical("数据库连接出错：", err)
		panic(err)
	}
	defer session.Close()

	c := session.DB(dbName).C(collectionName)
	err = c.Update(selector, change)

	return
}

func queryOpt(optType int, dbname string, collectionName string, query interface{}, result interface{}, sort string, skip, limit int) (count int, err error) {

	session, err := getDBConnSession()
	if err != nil {
		beego.Critical("数据库连接出错：", err)
		panic(err)
	}
	defer session.Close()

	q := session.DB(dbname).C(collectionName).Find(query)

	switch optType {
	case findOneOpt:
		err = q.One(result)
	case findListOpt:
		count, err = q.Count()
		err = q.Sort(sort).Skip(skip).Limit(limit).All(result)
	}
	return
}

//获取数据库连接
func getDBConnSession() (session *mgo.Session, err error) {
	session, err = mgo.Dial(conf.ConnAddr)
	return
}
