package models

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

var redisCache cache.Cache

func GetRedisClient() cache.Cache {
	if redisCache == nil {
		redis, err := cache.NewCache("redis", `{"conn":"`+beego.AppConfig.String("redisconn")+`", "key":"beecacheRedis"}`)
		if err != nil {
			log.Println(err, redis)
		}
		redisCache = redis
	}
	return redisCache
}
