package util

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
)



func NewModel() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@/bili?charset=utf8&parseTime=True&loc=Local")
	if err !=nil{
		fmt.Println("Connet Mysql DB error:",err.Error())
	}
	db.LogMode(true)
	db.DB().SetMaxOpenConns(100)   //设置数据库连接池最大连接数
	db.DB().SetMaxIdleConns(20)   //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	return  db
}