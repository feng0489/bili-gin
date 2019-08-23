package models

import (
	"bili-gin/entitys"
	"bili-gin/util"
	"log"
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

func FindUserById(id int64)  *entitys.User {
	db :=util.GetInstance().MyDB()
	var user entitys.User
	db.Select("id,username,nickname,phone,head_url,b_coin,user_tab,channel,store_up,focus,follower,create_time,last_time,last_ip").Where("id=?",id).First(&user)

    return &user
}
