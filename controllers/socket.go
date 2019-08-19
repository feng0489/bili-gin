package controllers

import (
	"bili-gin/util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}

type SocketMessage struct {
	MsgType int64 `json:"msg_type"`
	Code int64 `json:"code"`
	Uuid string `json:"uuid"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

//webSocket请求ping 返回pong
func Connet(c *gin.Context) {
	ip := util.GetIP(c)

	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		return
	}
	//defer ws.Close()
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()

		log.Println("Socket Message:",string(message))
		if err != nil {
			message = []byte(fmt.Sprintf("socket get message error:%v",err.Error()))
		}


		if string(message) == "ping" {
			message = []byte("pong")
			//写入ws数据
			err = ws.WriteMessage(mt, message)
			if err != nil {
				log.Println("send message error",err.Error())
			}
			break
		}
		if string(message) == "uuid" {
			uuids := util.GetCashe(strconv.Itoa(util.StringIpToInt(ip)))

			fmt.Println("cashe uuid:",uuids)
			var uuid string
			if uuids == nil{
				uuid = util.GetUuid()
				util.SetCashe(strconv.Itoa(util.StringIpToInt(ip)),uuid)
			}else{
				json.Unmarshal(uuids.([]byte), &uuid)
			}

			//写入ws数据
			err = ws.WriteMessage(mt, []byte(uuid))
			if err != nil {
				log.Println("send message error",err.Error())
			}
			break
		}
        var msg SocketMessage
		json.Unmarshal(message, &msg)
        if msg.Code == 0{
			msg.Code=102405
			msg.Msg = "Error Code"
			message = []byte(util.JsonEncode(msg))
			//写入ws数据
			err = ws.WriteMessage(mt, message)
			if err != nil {
				log.Println("send message error",err.Error())
			}
			break
		}
		//写入ws数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}
