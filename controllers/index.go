package controllers

import (
	"bili-gin/models"
	"bili-gin/util"
	"github.com/gin-gonic/gin"
)



func Index(c *gin.Context)  {
	val,_ := c.Cookie("sid")
	c.SetCookie("sid", "asdasdasdasdqweqwe", 3600, "/index", c.ClientIP(), false, false)
	c.JSON(200, gin.H{
		"msg": "ojbk",
		"Ip":util.GetIP(c),
		"cookie":val,
		"ClientIP":c.ClientIP(),
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