package models

import (
	"bili-gin/entitys"
	"bili-gin/util"
	"encoding/json"
	"fmt"
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
		fmt.Println("getCashe:",jsonStr)
		json.Unmarshal(jsonStr.([]byte), &all)
		fmt.Println("cashe:",all)
	}

	return &all
}