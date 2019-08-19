package controllers

import (
	"bili-gin/models"
	"bili-gin/util"
	"github.com/gin-gonic/gin"
)



func Index(c *gin.Context)  {

	c.JSON(200, gin.H{
		"msg": "ok",
		"Ip":util.GetIP(c),
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

}