package controllers

import (
	"bili-gin/models"
	"github.com/gin-gonic/gin"
)

func CarTab(c *gin.Context)  {
    tab:=c.Query("tab")

    carTab := models.GetCarByTab(tab)
	c.JSON(200, gin.H{
		"msg": "msgok",
		"code":200,
 		"data":carTab,
	})

}
