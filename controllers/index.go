package controllers

import (
	"bili-gin/models"
	"github.com/gin-gonic/gin"
)



func Index(c *gin.Context)  {

	c.JSON(200, gin.H{
		"msg": "msgok",
		"code":200,

	})
}

func TopNav(c *gin.Context)  {

	c.JSON(200, gin.H{
		"code":200,
		"msg": "ok",
		"data":models.TopMenu(),

	})
}

func HeadNav(c *gin.Context){
	c.JSON(200, gin.H{
		"code":200,
		"msg": "ok",
		"data":models.HeadMenu(),

	})
}