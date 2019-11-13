package main

import (
	"bili-gin/routers"
	"bili-gin/util"
	"log"
	"os"

)


func main() {


	//初始化mysql
	issucc := util.GetInstance().InitDataPool()
	if !issucc {
		log.Println("init database pool failure...")
		os.Exit(1)
	}
    //初始化redis
	issredis :=util.GetInredis().RedisPollInit()
	if !issredis{
		log.Println("connet redis poll failure")
	}
	//跑动一个http服务
    routers.Run()

}
