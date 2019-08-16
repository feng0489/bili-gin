package util

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"sync"
)

type MyRedisPool struct {

}


var (
	inredis *MyRedisPool
	onceinclude sync.Once
	redispool *redis.Pool
	err error
)



func GetInredis() *MyRedisPool {
	onceinclude.Do(func() {
		inredis = &MyRedisPool{}
	})
	return inredis
}



func (r *MyRedisPool)RedisPollInit() bool {

	redispool = &redis.Pool{
		MaxIdle:   3,
		MaxActive: 5,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				log.Fatal(err)
				return nil, err
			}

			//_, err = c.Do("AUTH", "123456")
			//if err != nil {
			//	log.Fatal(err)
			//	c.Close()
			//	return nil, err
			//}
			return c, err
		},
	}
	if redispool == nil {
		return true
	}else{
		return  false
	}


}



/*
* @fuc  对外获取数据库连接对象db
*/
func (m *MyRedisPool) MyRedis() (myredis redis.Conn) {
	return redispool.Get()
}


func SetCashe(key string,data interface{})(datas interface{},errors error){

	conn := redispool.Get()
	defer conn.Close()

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("jsonBytes:\n",string(jsonBytes))
	conn.Do("SET",key,string(jsonBytes))

	return data,nil
}


func GetCashe(key string) interface{}{
	conn := redispool.Get()
	defer conn.Close()

    jsonStr,err := conn.Do("GET",key)

    if err != nil {
		log.Println(err)
	}
	return jsonStr
}