package main

import (
	"fmt"
	"net"
	"time"
	"gotcp/chatroom/server/model"
)

func init(){
	// 当服务器开启后就初始化连接池
	initPool("localhost:6379",8,0,180 * time.Second)
	initUserDao()
}

// 编写一个函数完成对UserDao的初始化
func initUserDao(){
	// 这里要注意一个初始化顺序问题
	// 先调用initPool 再调用 initUserDao
	model.MyUserDao = model.NewUserDao(pool) // 这里的pool来源于redis中的pool
}

func process(conn net.Conn){
	defer conn.Close()
	// 循环读取
	processall := &Processall{
		Conn : conn,
	}
	processall.processalltwo()
}

func main(){
	listener,err := net.Listen("tcp","0.0.0.0:8889")
	defer listener.Close()
	if err != nil{
		fmt.Println("net.Listener err = ",err)
		return
	}

	fmt.Printf("服务器网络类型为:%v 地址为:%v\n",listener.Addr().Network(),listener.Addr().String())
	for{
		// Accept返回listen服务端的连接并等待下一个连接
		conn,err := listener.Accept()
		defer conn.Close()
		if err != nil{
			fmt.Println("listener.Accept err = ",err)
			return
		}

		fmt.Println("客户端连接成功！ip地址为:",conn.RemoteAddr())
		go process(conn)
	}
}