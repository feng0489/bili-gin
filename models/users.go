package models

import (
	"bili-gin/entitys"
	"bili-gin/util"
	"log"
	"fmt"
)

func Register(Users *entitys.Users) int64 {

	db :=util.GetInstance().MyDB()
	res:=db.Create(Users)
	if res.Error != nil {
		log.Println("insert users error:",res.Error.Error())
		return 0
	}else{
		return Users.Id
	}

}

func CheckReg(re *entitys.UserReg) bool {
	var ok bool=false
	db :=util.GetInstance().MyDB()
	var count entitys.Result
    var Users entitys.Users
	sql := fmt.Sprintf("select count(id) as total from bili_users where (nickname='%v') or (phone='%v') ",re.NickName,re.Phone)

	db.Table(Users.TableName()).Raw(sql).Select("total").Scan(&count)
	log.Println("conut result total:",count.Total)
	if count.Total > 0{
		ok= true
	}
	return ok
}

func CheckNickName(nick string) bool{
	var ok bool=false
	db :=util.GetInstance().MyDB()
	var count entitys.Result
	var Users entitys.Users
	sql := fmt.Sprintf("select count(id) as total from bili_users where nickname='%v' ",nick)
	db.Table(Users.TableName()).Raw(sql).Select("total").Scan(&count)
	if count.Total > 0{
		ok= true
	}
	return ok
}


func CheckPhone(phone string) bool{
	var ok bool=false
	db :=util.GetInstance().MyDB()
	var count entitys.Result
	var Users entitys.Users
	sql := fmt.Sprintf("select count(id) as total from bili_users where phone='%v' ",phone)
	db.Table(Users.TableName()).Raw(sql).Select("total").Scan(&count)
	if count.Total > 0{
		ok= true
	}
	return ok
}


func FindUserById(id int64)  *entitys.Users {
	db :=util.GetInstance().MyDB()
	var Users entitys.Users
	res :=db.Table(Users.TableName()).Select("id,username,nickname,phone,head_url,b_coin,user_tab,channel,store_up,focus,follower,create_time,last_time,last_ip").Where("id=?",id).First(&Users)
    if res.Error != nil {
    	log.Println("users=>models=>FindUserById:",res.Error.Error())
	}
    return &Users
}


func FindUserByPhone(phone string) *entitys.Users {
	db :=util.GetInstance().MyDB()
	var Users entitys.Users
	res := db.Select("id,password,username,nickname,phone,head_url,b_coin,user_tab,channel,store_up,focus,follower").Where("phone=?",phone).First(&Users)
	if res.Error != nil {
		log.Println("FindUserByPhone error :" ,res.Error.Error())
		return  nil
	}
	return &Users
}


func UpUser(id int64,data map[string]interface{}) bool {
	db :=util.GetInstance().MyDB()
	var Users entitys.Users
	res :=db.Table(Users.TableName()).Where("id=?", id).Updates(data)
	if res.Error != nil{
		log.Println("UpUser",res.Error.Error())
		return false
	}else{
		return true
	}
}