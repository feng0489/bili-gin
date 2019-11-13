package util

import (
	"bili-gin/entitys"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"log"
	"fmt"
	"math/rand"
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

func SetUserCashe(c *gin.Context,id int64,user *entitys.Users) interface{} {
	key := fmt.Sprintf("%v_user_%v",StringIpToInt(GetIP(c)),id)
	u,_ := SetCashe(key,user)
	return u
}

func GetUserCashe(c *gin.Context,id int64) (user *entitys.Users) {
	key := fmt.Sprintf("%v_user_%v",StringIpToInt(GetIP(c)),id)
	users := GetCashe(key)
	if users == nil {
		return nil
	}
	json.Unmarshal(users.([]byte), user)
	return user
}

func GetPostId(c *gin.Context)int64{
	ids:=c.Request.FormValue("id")
	if ids == "" {

		return 0
	}
	id, err := strconv.ParseInt(ids,10,64)
	if err != nil {
		log.Println("GetPostId error:", err.Error())
		return 0
	}
	return  id
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
			if nowtime - info.CreateTime > 144000{
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
	SetCashe(fmt.Sprintf("%v_%v",StringIpToInt(GetIP(c)),key),cookie)
	c.SetCookie(key, value, 30*1000, "/", "localhost", false, true)
}

func GetCookie(c *gin.Context,key string)string{
	cookie, err := c.Request.Cookie(fmt.Sprintf("%v_%v",StringIpToInt(GetIP(c)),key))
	if err != nil {
		return ""
	}else{
		return cookie.Value
	}
}

func CheckCookie(c *gin.Context,key string) bool {
	reqCookie :=GetCookie(c,key)
	var casheCookie http.Cookie
	casheCookies := GetCashe(fmt.Sprintf("%v_%v",StringIpToInt(GetIP(c)),key))
	json.Unmarshal(casheCookies.([]byte), &casheCookie)
	if reqCookie == casheCookie.Value {
		return true
	}else{
		return false
	}

}

func GetCacheToken(c *gin.Context,key string) string{
	var casheCookie http.Cookie
	casheCookies := GetCashe(fmt.Sprintf("%v_%v",StringIpToInt(GetIP(c)),key))
	json.Unmarshal(casheCookies.([]byte), &casheCookie)
	return casheCookie.Value
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

func CreateUserName(hpone string) string{
	nameLast := hpone[len(hpone)-6:]
	return fmt.Sprintf("%v%v",RandNum(1),nameLast)
}

func Md5(pass string) string{
	h := md5.New()
	h.Write([]byte(pass)) // 需要加密的字符串为 password
	submit :=hex.EncodeToString(h.Sum(nil))
	return submit
}


func RandNum(width int) string {
	numeric := [10]byte{1,2,3,4,5,6,7,8,9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[ rand.Intn(r) ])
	}
	return sb.String()
}

func CreateToken(c *gin.Context,id int64) string {
	one :=RandNum(6)
	two := fmt.Sprintf("%v%v",RandNum(3),id)
	tree := time.Now().Unix()
	four := StringIpToInt(GetIP(c))

	return fmt.Sprintf("%v_%v_%v_%v",one,two,tree,four)

}