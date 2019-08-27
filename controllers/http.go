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
		r:=c.Request
		// 处理js-ajax跨域问题
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Add("Access-Control-Allow-Headers", "Access-Token")

		ip :=util.GetIP(c)
		path :=c.Request.URL.Path
		uuid:=c.Request.FormValue("uuid")


		if path== "/favicon.ico" {
			c.Abort()
			return
		}


		if  isCheckToken(path) && util.GetCookie(c,"token") == ""{
			c.JSON(504, gin.H{
				"code":504,
				"msg": "bad request did",
				"url":c.Request.URL.Path,
			})
			c.Abort()
			return
		}

		if  isCheckToken(path) && !util.CheckCookie(c,"token"){
			c.JSON(504, gin.H{
				"code":504,
				"msg": "bad request done",
				"url":c.Request.URL.Path,
			})
			c.Abort()
			return
		}



		uid := util.GetUuid(ip)

		if checkoutUuid(path) && uuid == "" {
			c.JSON(504, gin.H{
				"code":504,
				"msg": "bad request path",
				//"url":c.Request.URL.Path,
			})
			c.Abort()
			return
		}
		if checkoutUuid(path) && uuid != uid {
			c.JSON(505, gin.H{
				"code":505,
				"msg": "bad request got",
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

	checkout := [4]string{"/userinfo"}
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

