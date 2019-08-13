package main

import (
	"bili-gin/routers"
	"fmt"
	"github.com/jinzhu/gorm"
)

var MyDB *gorm.DB
func init(){
	MyDB, err := gorm.Open("mysql", "root:root@/bili?charset=utf8&parseTime=True&loc=Local")
	if err !=nil{
		fmt.Println("Connet Mysql DB error:",err.Error())
	}
	MyDB.DB().SetMaxOpenConns(100)   //设置数据库连接池最大连接数
	MyDB.DB().SetMaxIdleConns(20)   //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
}
func main() {
	MyDB.LogMode(true)
    routers.Run()
}

