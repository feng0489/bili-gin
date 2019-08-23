package controllers

import (
	"bili-gin/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func FindUsers(c *gin.Context)  {
	ids:=c.Query("id")
	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil {
		c.JSON(200, gin.H{
			"msg": err.Error(),
			"code":12304,
			"data":"",
		})
		return
	}
	if id <=0 {
		c.JSON(200, gin.H{
			"msg": "msg error",
			"code":12304,
			"data":"",
		})
		return
	}

	user := models.FindUserById(id)

	c.JSON(200, gin.H{
		"msg": "msg ok",
		"code":200,
		"data":user,
	})
}
