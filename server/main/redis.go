package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

/*
	初始化工作
*/
var pool *redis.Pool

// time.Duration ,Duration 类型代表两个时间点之间经过的时间，ns为单位
func initPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) { // 定义成 initPool 而不是 init ，所以不会自动调用
	/*
		为什么要定义成 initPool ，目的就是设置 MaxIdel 等一系列的参数
	*/
	//pool := &redis.Pool{
	//	MaxIdle:         8,
	//	MaxActive:       0,   //表示和数据库最大的连接数，0表示没有限制
	//	IdleTimeout:     100, //最大空闲时间
	//	MaxConnLifetime: 0,
	//	Dial: func() (redis.Conn, error) {
	//		return redis.Dial("tcp", "localhost¬:6379")
	//	},
	//}
	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
	}
}
