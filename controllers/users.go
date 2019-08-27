package controllers

import (
	"bili-gin/entitys"
	"bili-gin/models"
	"bili-gin/util"
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"strconv"
	"time"
)



func UserRegister(c *gin.Context)  {
	var user entitys.UserReg
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error(),"code":12304,})
		return
	}
    if models.CheckReg(&user) {
		c.JSON(200, gin.H{
			"msg": "this nick_name or phone is already used",
			"code":12305,
			"data":"",
		})
		return
	}
	newTime :=time.Now().Unix()


	newUser := new(entitys.User)
	newUser.Username =util.CreateUserName(user.Phone)
	newUser.Phone = user.Phone
	newUser.CreateTime = newTime
	newUser.Nickname = user.NickName
	newUser.Password = util.Md5(user.Password)
	id:= models.Register(newUser)

	if id > 0 {
		util.SetCookie(c,"token",util.GetToken(c,id))
		c.JSON(200, gin.H{
			"msg": "ok",
			"code":200,
			"data":id,
		})
	}else{
		c.JSON(200, gin.H{
			"msg": "register error please try again",
			"code":12306,
			"data":0,
		})
	}

}

func CheckNickName(c *gin.Context){
	nick := c.DefaultPostForm("nick_name", "") // 此方法可以设置默认值
	if nick == ""{
		c.JSON(200, gin.H{
			"msg": "empty nick",
			"code":12304,
		})
		return
	}
	if models.CheckNickName(nick) {
		c.JSON(200, gin.H{
			"msg": fmt.Sprintf("this %v is already used",nick),
			"code":12305,
		})
	}else{
		c.JSON(200, gin.H{
			"msg":   fmt.Sprintf("this %v can be use",nick),
			"code":200,
		})
	}
}

func CheckPhone(c *gin.Context){
	phone := c.DefaultPostForm("phone", "") // 此方法可以设置默认值
	if phone == ""{
		c.JSON(200, gin.H{
			"msg": "empty phone",
			"code":12304,
		})
		return
	}

	if models.CheckPhone(phone) {
		c.JSON(200, gin.H{
			"msg":  fmt.Sprintf("this %v is already used",phone),
			"code":12305,
		})
	}else{
		c.JSON(200, gin.H{
			"msg": fmt.Sprintf("this %v can be use",phone),
			"code":200,
		})
	}

}

func UserLogin(c *gin.Context)  {
	phone := c.DefaultPostForm("phone", "") // 此方法可以设置默认值
	pass := c.DefaultPostForm("pass", "") // 此方法可以设置默认值
	if phone == "" || pass == ""{
		c.JSON(200, gin.H{
			"msg": "request error",
			"code":12304,
		})
		return
	}

	user:=models.FindUserByPhone(phone)
	if user == nil {
		c.JSON(200, gin.H{
			"msg": "username or password error",
			"code":12305,
		})
		return
	}
	if user.Id ==0 {
		c.JSON(200, gin.H{
			"msg": "username or password error",
			"code":12306,
		})
		return
	}
	submit := util.Md5(pass)
	if submit != user.Password{
		c.JSON(200, gin.H{
			"msg": "username or password error",
			"code":12307,
			"pass":submit,
		})
		return
	}
	Updata :=map[string]interface{}{
        "last_time":time.Now().Unix(),
        "last_ip":util.StringIpToInt(util.GetIP(c)),
	}
	token :=util.GetToken(c,user.Id)
	util.SetCookie(c,"token",token)
	models.UpUser(user.Id,Updata)

	c.JSON(200, gin.H{
		"msg": "msgok",
		"code":200,
		"token":token,
		"data":map[string]interface{}{
			"id":user.Id,
			"nickname":user.Nickname,
			"head_url":user.HeadUrl,
		},
	})



}

func FindUsers(c *gin.Context)  {
	ids:=c.Query("id")
	id, err := strconv.ParseInt(ids, 10, 64)
	if err != nil {
		c.JSON(200, gin.H{
			"msg": err.Error(),
			"code":12304,
			"data":"",
		})
		return
	}
	if id <=0 {
		c.JSON(200, gin.H{
			"msg": "msg error",
			"code":12304,
			"data":"",
		})
		return
	}

	user := models.FindUserById(id)

	c.JSON(200, gin.H{
		"msg": "msg ok",
		"code":200,
		"data":user,
	})
}
