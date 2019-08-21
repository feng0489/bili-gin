package util

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"log"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type UuidInfo struct {
	Uuid string `json:"uuid"`
	CreateTime int64 `json:"create_time"`
}

func JsonEncode(data interface{}) string {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	return string(jsonBytes)
}


func GetIP(c *gin.Context) string{
	ip := c.ClientIP()
	if ip == ""{
		ips:=strings.Split(c.Request.RemoteAddr, ":")
		ip= ips[0]
	}
	return ip
}


func GetUuid(ip string) string {
	if ip == ""{
		id := uuid.NewV4()
		return fmt.Sprintf("%v",id)
	}else{
		uuids := GetCashe(fmt.Sprintf("%v_uuid",ip))
		var info UuidInfo
		 nowtime :=time.Now().Unix()
		if uuids != nil{
			json.Unmarshal(uuids.([]byte), &info)
			if nowtime - info.CreateTime > 100{
				info.Uuid = fmt.Sprintf("%v",uuid.NewV4())
				info.CreateTime = time.Now().Unix()
				SetCashe(fmt.Sprintf("%v_uuid",ip),info)
				return info.Uuid
			}else{
				return info.Uuid
			}

		}else{
			info.Uuid = fmt.Sprintf("%v",uuid.NewV4())
            info.CreateTime = time.Now().Unix()
			SetCashe(fmt.Sprintf("%v_uuid",ip),info)
            return info.Uuid

		}

	}

}

func SetCookie(c *gin.Context,key,value string)  {
	cookie := &http.Cookie{
		Name: key,
		Value:value,
		Domain:c.Request.Host,
		Expires:time.Now().AddDate(0, 0, 1),
		MaxAge:1000*36,
		HttpOnly:false,
	}
	SetCashe(fmt.Sprintf("%v_%v",GetIP(c),key),cookie)
	http.SetCookie(c.Writer, cookie)
}

func GetCookie(c *gin.Context,key string)string{
	cookie, err := c.Request.Cookie(fmt.Sprintf("%v_%v",GetIP(c),key))
	if err != nil {
		return ""
	}else{
		return cookie.Value
	}
}

func CheckCookie(c *gin.Context,key string) bool {
	reqCookie :=GetCookie(c,key)
	var casheCookie http.Cookie
    casheCookies := GetCashe(fmt.Sprintf("%v_%v",GetIP(c),key))
	json.Unmarshal(casheCookies.([]byte), &casheCookie)
    if reqCookie == casheCookie.Value {
		return true
	}else{
		return false
	}
	
}


func StringIpToInt(ipstring string) int {
	ipSegs := strings.Split(ipstring, ".")
	var ipInt int = 0
	var pos uint = 24
	for _, ipSeg := range ipSegs {
		tempInt, _ := strconv.Atoi(ipSeg)
		tempInt = tempInt << pos
		ipInt = ipInt | tempInt
		pos -= 8
	}
	return ipInt
}


func IpIntToString(ipInt int) string {
	ipSegs := make([]string, 4)
	var len int = len(ipSegs)
	buffer := bytes.NewBufferString("")
	for i := 0; i < len; i++ {
		tempInt := ipInt & 0xFF
		ipSegs[len-i-1] = strconv.Itoa(tempInt)
		ipInt = ipInt >> 8
	}
	for i := 0; i < len; i++ {
		buffer.WriteString(ipSegs[i])
		if i < len-1 {
			buffer.WriteString(".")
		}
	}
	return buffer.String()
}