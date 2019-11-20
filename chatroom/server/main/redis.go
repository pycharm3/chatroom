// 初始化pool连接池
package main

import (
	"time"
	"github.com/garyburd/redigo/redis"
)

// 声明一个redis.Pool结构体类型的pool
// 定义一个全局的pool
var pool *redis.Pool

// 程序启动初始化连接池
func initPool(address string,maxIdle int,maxActive int,idleTimeout time.Duration){
	pool = &redis.Pool{
		// 最大空闲连接数
		MaxIdle : maxIdle,
		// 表示和数据库的最大连接数，0表示不限制
		MaxActive : maxActive,
		// 最大空闲时间
		IdleTimeout : idleTimeout,
		// 初始化连接，表示连接到哪个地址
		Dial : func() (redis.Conn, error){
			return redis.Dial("tcp",address)
		},
	}
}