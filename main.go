package main

import (
	"bili-gin/routers"
	"bili-gin/util"
	"log"
	"os"
)


func main() {

	issucc := util.GetInstance().InitDataPool()
	if !issucc {
		log.Println("init database pool failure...")
		os.Exit(1)
	}

	issredis :=util.GetInredis().RedisPollInit()
	if !issredis{
		log.Println("connet redis poll failure")
	}
    routers.Run()
}

