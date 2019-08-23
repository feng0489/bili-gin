package routers

import (
	"bili-gin/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Run(){

	//gin初始化
	r := gin.Default()
	r.Use(controllers.CheckHttp())
	r.GET("/index", controllers.Index)

    //nav
	r.GET("/top", controllers.TopNav)
	r.GET("/head", controllers.HeadNav)


	//carousel
	r.GET("/tab",controllers.CarTab)

	//users
	r.GET("/user",controllers.FindUsers)
	r.POST("/reg",controllers.UserRegister)

	//socket
	r.GET("/connet", controllers.Connet)




    //http 配置
	s := &http.Server{
		Addr:           ":6619",//监听端口
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}



	fmt.Println("Golang Going To Run A Http Server Port :",s.Addr)


	err := s.ListenAndServe()//启动服务
	if err != nil {
		fmt.Println("bili-gin run Server error:",err.Error())
	}

}


