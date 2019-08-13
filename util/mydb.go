package util

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	)

func NewMysql() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@/bili?charset=utf8&parseTime=True&loc=Local")
	if err !=nil{
       fmt.Println("Connet Mysql DB error:",err.Error())
	}
	db.LogMode(true)
	return db
}