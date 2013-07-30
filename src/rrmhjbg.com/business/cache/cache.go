package cache

import (
	"github.com/astaxie/beego"
)

type Cache struct {
	resCache *beego.BeeCache
}

func (this *Cache) Init() {
	this.resCache = beego.NewBeeCache()
	this.resCache.Every = 0
	this.resCache.Start()
}

func (this *Cache) Put(key string, val interface{}) {
	err := this.resCache.Put(key, val, 0)
	if err != nil {
		beego.Error(err)
	}
}

func (this *Cache) IsExist(key string) bool {
	return this.resCache.IsExist(key)
}

func (this *Cache) Get(key string) (val interface{}) {
	return this.resCache.Get(key)
}

func (this *Cache) Del(key string) {
	_, err := this.resCache.Delete(key)
	if err != nil {
		beego.Error(err)
	}
}
