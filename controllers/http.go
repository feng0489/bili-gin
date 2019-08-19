package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func CheckHttp() gin.HandlerFunc{
	return func(c *gin.Context){
		fmt.Println("befor request>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

		path :=c.Request.URL.Path
		token:=c.Request.FormValue("api_token")
		uuid:=c.Request.FormValue("uuid")


		if  checkoutToken(path) && token == ""{
			c.JSON(404, gin.H{
				"code":404,
				"msg": "can't request without token",
				"url":c.Request.URL.Path,
			})
			c.Abort()
			return
		}

		if checkoutUuid(path) && uuid == "" {

		}

		c.Next()
	}

}

/**
 * 过滤不需要校验token的请求
 */
func checkoutToken(url string) bool {

	checkout := [3]string{"/index","/top","/connet"}
	var check bool=true
	for _,v:=range checkout{
		if v==url{
			check=false
		}
	}
	return check
}
/**
 * 过滤不需要校验uuid的请求
 */
func checkoutUuid(url string) bool {
	checkout := [3]string{"/index","/top","/connet"}
	var check bool=true
	for _,v:=range checkout{
		if v==url{
			check=false
		}
	}
	return check
}

