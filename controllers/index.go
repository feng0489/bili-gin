package controllers

import (
	"bili-gin/models"
	"bili-gin/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)



func Index(c *gin.Context)  {
	ip :=util.GetIP(c)
	v :=strconv.FormatInt(time.Now().Unix(),10)
	cookie := &http.Cookie{
		Name:  "token",
		Value:v,
	}

	http.SetCookie(c.Writer, cookie)
	c.JSON(200, gin.H{
		"msg": "ojbk",
		"Ip":ip,
		"cookie":v,
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