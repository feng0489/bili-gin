package controllers

import (
	"bili-gin/util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckHttp() gin.HandlerFunc{
	return func(c *gin.Context){
		fmt.Println("befor request>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		path :=c.Request.URL.Path
		fmt.Println("path>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>:",path	)

		token:=c.Request.FormValue("api_token")
		uuid:=c.Request.FormValue("uuid")
        reqType :=c.Request.Header.Get("Upgrade")
        ip :=util.GetIP(c)
		if  checkoutToken(path) && token == ""{
			c.JSON(504, gin.H{
				"code":504,
				"msg": "can't request without token",
				"url":c.Request.URL.Path,
			})
			c.Abort()
			return
		}

		if checkoutUuid(path) && uuid == "" {
			c.JSON(504, gin.H{
				"code":504,
				"msg": "can't request without uuid",
				"url":c.Request.URL.Path,
			})
			c.Abort()
			return
		}
		if checkoutCookie(path) {

			if cookie, err := c.Request.Cookie("sid"); err == nil  {
				if len(cookie.Value)==0 {
					c.JSON(504, gin.H{
						"code":504,
						"msg": "request error",
						"url":c.Request.URL.Path,
					})
					c.Abort()
					return
				}

				var cookieCashe string
				ckcashe := util.GetCashe(fmt.Sprintf("%v_sid",ip))
				json.Unmarshal(ckcashe.([]byte), &cookieCashe)
				if cookie.Value != cookieCashe{
					c.JSON(504, gin.H{
						"code":504,
						"msg": "request with a error",
						"url":c.Request.URL.Path,
					})
					c.Abort()
					return
				}


			}

		}

		uuids := util.GetCashe(ip)
		if checkoutUuid(path) && uuids == nil{
			c.JSON(504, gin.H{
				"code":504,
				"msg": "error uuid",
				"url":c.Request.URL.Path,
			})
			c.Abort()
			return
		}

		if checkoutUuid(path) && uuids != nil{
			var uid string
			json.Unmarshal(uuids.([]byte), &uid)
			if checkoutUuid(path) && uuids == uid {
				c.JSON(504, gin.H{
					"code":504,
					"msg": "can't request with a error uuid",
					"url":c.Request.URL.Path,
				})
				c.Abort()
				return
			}
		}

        if  reqType == "websocket" {
			key :=c.Request.Header.Get("Sec-Websocket-Key")

			if cookie, err := c.Request.Cookie("sid"); err == nil {
				util.SetCashe(fmt.Sprintf("%v_sid",ip),key)
				cookie.Value = key
				http.SetCookie(c.Writer, cookie)
			}
		}

		c.Next()
	}

}

/**
 * 过滤不需要校验token的请求
 */
func checkoutToken(url string) bool {

	checkout := [5]string{"/index","/top","/connet","/ws","/allProject"}
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
	checkout := [5]string{"/index","/top","/connet","/ws","/allProject"}
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