package models

import (
	"bili-gin/entitys"
	"bili-gin/util"
	"log"
	"fmt"
)

func Register(user *entitys.User) int64 {

	db :=util.GetInstance().MyDB()
	res:=db.Create(user)
	if res.Error != nil {
		log.Println("insert users error:",res.Error.Error())
		return 0
	}else{
		return user.Id
	}

}

func CheckReg(re *entitys.UserReg) bool {
	var ok bool=false
	db :=util.GetInstance().MyDB()
	var count entitys.Result
    var user entitys.User
	sql := fmt.Sprintf("select count(id) as total from bili_users where (nickname='%v') or (phone='%v') ",re.NickName,re.Phone)

	db.Table(user.TableName()).Raw(sql).Select("total").Scan(&count)
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
	var user entitys.User
	sql := fmt.Sprintf("select count(id) as total from bili_users where nickname='%v' ",nick)
	db.Table(user.TableName()).Raw(sql).Select("total").Scan(&count)
	if count.Total > 0{
		ok= true
	}
	return ok
}


func CheckPhone(phone string) bool{
	var ok bool=false
	db :=util.GetInstance().MyDB()
	var count entitys.Result
	var user entitys.User
	sql := fmt.Sprintf("select count(id) as total from bili_users where phone='%v' ",phone)
	db.Table(user.TableName()).Raw(sql).Select("total").Scan(&count)
	if count.Total > 0{
		ok= true
	}
	return ok
}


func FindUserById(id int64)  *entitys.User {
	db :=util.GetInstance().MyDB()
	var user entitys.User
	db.Select("id,username,nickname,phone,head_url,b_coin,user_tab,channel,store_up,focus,follower,create_time,last_time,last_ip").Where("id=?",id).First(&user)

    return &user
}


func FindUserByPhone(phone string) *entitys.User {
	db :=util.GetInstance().MyDB()
	var user entitys.User
	res := db.Select("id,password,username,nickname,phone,head_url,b_coin,user_tab,channel,store_up,focus,follower").Where("phone=?",phone).First(&user)
	if res.Error != nil {
		log.Println("FindUserByPhone error :" ,res.Error.Error())
		return  nil
	}
	return &user
}


func UpUser(id int64,data map[string]interface{}) bool {
	db :=util.GetInstance().MyDB()
	var user entitys.User
	res :=db.Table(user.TableName()).Where("id=?", id).Updates(data)
	if res.Error != nil{
		log.Println("UpUser",res.Error.Error())
		return false
	}else{
		return true
	}
}