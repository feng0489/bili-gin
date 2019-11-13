package controllers

import (
	"bili-gin/entitys"
	"bili-gin/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AllProjeck(c *gin.Context)  {

	db :=util.GetInstance().MyDB()
	var deel entitys.Deel
	var deels []entitys.Deel
	var times = make(chan int64)
	db.Table(deel.TableName()).Select("repay_start_time ").Where("deal_status>=0 and publish_wait=0 and is_effect=1 and is_delete=0").Group("FROM_UNIXTIME(repay_start_time,'%Y-%m-%d')").Find(&deels)

	go func() {
		for _,d:=range deels {
			times <-d.RepayStartTime
		}
		close(times)
	}()
	fmt.Println(<-times)

	deelData :=make([]map[string]interface{},len(deels))
	c.JSON(200, gin.H{
		"code":200,
		"msg": "ok",
		"data":deelData,

	})
}