package controllers

import (
	"bili-gin/entitys"
	"bili-gin/util"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context)  {

	c.JSON(200, gin.H{
		"msg": "ok",
	})
}

func TopNav(c *gin.Context)  {
	db :=util.NewMysql()
	var all []entitys.TopNav
	db.Find(&all)
	c.JSON(200, gin.H{
		"code":200,
		"msg": "ok",
		"data":all,
	})
}