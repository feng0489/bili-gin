package util

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"log"
	"fmt"
	"encoding/json"
	"strconv"
	"strings"

)

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


func GetUuid() string {
	id,_:= uuid.NewV4()
	return fmt.Sprintf("%v",id)
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