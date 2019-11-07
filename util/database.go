package util

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"sync"
)


/*
* MysqlConnectiPool
* 数据库连接操作库
* 基于gorm封装开发
 */
type MysqlConnectiPool struct {
}


var (
 instance *MysqlConnectiPool
 once sync.Once
 db *gorm.DB
 err_db error
)



func GetInstance() *MysqlConnectiPool {
	once.Do(func() {
		instance = &MysqlConnectiPool{}
	})
	return instance
}


/*
* @fuc 初始化数据库连接(可在mail()适当位置调用)
*/
func (m *MysqlConnectiPool) InitDataPool() (issucc bool) {
	db, err_db = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/outsite_gxxew?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	if err_db != nil {
		log.Fatal(err_db)
		return false
	}
	//关闭数据库，db会被多个goroutine共享，可以不调用
	// defer db.Close()
	return true
}


/*
* @fuc  对外获取数据库连接对象db
*/
func (m *MysqlConnectiPool) MyDB() (db_con *gorm.DB) {
	return db
}
