package models

import (
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

var redisCache cache.Cache

func GetRedisClient() cache.Cache {
	if redisCache == nil {
		redisPwd := beego.AppConfig.String("redispwd")
		redisConn := `{"conn":"`+beego.AppConfig.String("redisconn")+`", "key":"beecacheRedis","password":"`+redisPwd+`"}`
		if redisPwd == ""{
			redisConn = `{"conn":"`+beego.AppConfig.String("redisconn")+`", "key":"beecacheRedis"}`
		}
		fmt.Println(redisConn)
		redis, err := cache.NewCache("redis", redisConn)
		if err != nil {
			log.Println(err, redis)
		}
		redisCache = redis
	}
	return redisCache
}
