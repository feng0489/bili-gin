package controllers

import (
	"bili-gin/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CheckHttp() gin.HandlerFunc{
	return func(c *gin.Context){
		fmt.Println("befor request>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")

		w := c.Writer
		// 处理js-ajax跨域问题
		w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Add("Access-Control-Allow-Headers", "Access-Token")

		ip :=util.GetIP(c)
		path :=c.Request.URL.Path
		uuid:=c.Request.FormValue("uuid")

		token:=c.Request.FormValue("api_token")
		if  isCheckToken(path) && token == ""{
			c.JSON(504, gin.H{
				"code":504,
				"msg": "bad request",
				//"url":c.Request.URL.Path,
			})
			c.Abort()
			return
		}


		uid := util.GetUuid(ip)

		if checkoutUuid(path) && uuid == "" {
			c.JSON(504, gin.H{
				"code":504,
				"msg": "bad request",
				//"url":c.Request.URL.Path,
			})
			c.Abort()
			return
		}
		if checkoutUuid(path) && uuid != uid {
			c.JSON(505, gin.H{
				"code":505,
				"msg": "bad request",
				//"url":c.Request.URL.Path,
			})
			c.Abort()
			return
		}

		c.Next()
	}

}

/**
 * 过滤需要校验token的请求
 */
func isCheckToken(url string) bool {

	checkout := [4]string{"/userInfo"}
	var check bool=false
	for _,v:=range checkout{
		if v==url{
			check=true
		}
	}
	return check
}
/**
 * 过滤不需要校验uuid的请求
 */
func checkoutUuid(url string) bool {
	checkout := [4]string{"/index","/connet","/ws"}
	var check bool=true
	for _,v:=range checkout{
		if v==url{
			check=false
		}
	}
	return check
}


/**
 * 过滤不需要检验cookie的请求
 */

func checkoutCookie(url string) bool {
	checkout := [4]string{"/index","/top","/connet","/ws"}
	var check bool=true
	for _,v:=range checkout{
		if v==url{
			check=false
		}
	}
	return check
}