package models

import (
	"bili-gin/entitys"
	"bili-gin/util"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
)

func TopMenu() *[]entitys.TopNav {
	var all []entitys.TopNav
	 jsonStr :=  util.GetCashe("TopNav")

	if jsonStr == nil {
		db :=util.GetInstance().MyDB()
		db.Select("id,name,action,icon").Find(&all)
		util.SetCashe("TopNav",&all)
	}else{
		json.Unmarshal(jsonStr.([]byte), &all)
	}

	return &all
}

func HeadMenu() *[]entitys.HeadNav{
	var all []entitys.HeadNav
	jsonStr :=  util.GetCashe("HeadNav")

	if jsonStr == nil {
		db :=util.GetInstance().MyDB()
		db.Select("id,name,action,icon,nav_id,new_count").Find(&all)

		util.SetCashe("HeadNav",&all)
	}else{
		json.Unmarshal(jsonStr.([]byte), &all)

	}
	return &all
}