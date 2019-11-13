package models

import (
	"bili-gin/entitys"
	"bili-gin/util"
)

func GetCarByTab(tab string) *[]entitys.Carousel{
    var cars []entitys.Carousel
	db :=util.GetInstance().MyDB()
	if tab == ""{
		db.Select("id,title,photo,jump_url,tab_name,tab").
			Where("opne=1").
			Order("sort").
			Find(&cars)
	}else{
		db.Select("id,title,photo,jump_url,tab_name,tab").
			Where("opne=1").
			Where("tab=?",tab).
			Order("sort").
			Find(&cars)
	}
    return  &cars
}

