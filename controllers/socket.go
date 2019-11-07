package controllers

import (
	"bili-gin/util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}


var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
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
		var msg SocketMessage
		json.Unmarshal(message, &msg)


		if err != nil {

			//log.Println("socket connet error")
			//message = []byte(fmt.Sprintf("socket get message error:%v",err.Error()))
			err = ws.WriteMessage(mt, []byte(err.Error()))
			if err != nil {
				log.Println("send message error",err.Error())
			}
			break
		}

		if msg.MsgType == 0 {
			msg.Code=102507
			msg.Msg="socket get msg_type error"
			message =[]byte(util.JsonEncode(msg))
			err = ws.WriteMessage(mt, message)
			if err != nil {
				log.Println("send message error",err.Error())
			}
			break
		}


		if msg.MsgType == 1001 {
			message = []byte("pong")

			msg.Code=200
			msg.Msg="pong"
			msg.MsgType=1001
			msg.Data =  map[string]interface{}{"pong": "pong"}
			message = []byte(util.JsonEncode(msg))
			//写入ws数据
			err = ws.WriteMessage(mt, message)
			if err != nil {
				log.Println("send message error",err.Error())
			}
			break
		}


		if msg.MsgType == 1002 {
			uuids := util.GetCashe(ip)

			fmt.Println("cashe uuid:",uuids)
			var uuid string
			if uuids == nil{
				uuid = util.GetUuid()
				util.SetCashe(ip,uuid)
			}else{
				json.Unmarshal(uuids.([]byte), &uuid)
			}

			if uuid == ""{
				msg.Code=102404
				msg.Msg = "Error uuid create nil"
				message = []byte(util.JsonEncode(msg))
				//写入ws数据
				err = ws.WriteMessage(mt,message)
				if err != nil {
					log.Println("send message error",err.Error())
				}
				break
			}

			msg.Code=200
			msg.Msg="ok"
			msg.MsgType=1002
			msg.Uuid = uuid
			msg.Data =  map[string]interface{}{"uuid": uuid}
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

func Wshandler(w http.ResponseWriter, r *http.Request)  {
	conn, err := wsupgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()

		log.Println("socket connet id:",t)
		if err != nil {
			break
		}

		conn.WriteMessage(t, msg)
	}
}

