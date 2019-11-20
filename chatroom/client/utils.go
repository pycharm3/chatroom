package main

import (
	"fmt"
	"net"
	"gotcp/chatroom/common/message"
	"encoding/binary"
	"encoding/json"
)

// 编写一个函数 readMsg 用来读取消息
func readMsg(conn net.Conn)(msg message.Message,err error){
	buf := make([]byte,1024)
	fmt.Println("准备读取客户端传输数据...")
	// 从conn中读取4个字节放入bug中
	_,err = conn.Read(buf[:4])
	if err != nil{
		// fmt.Println("conn.Read err = ",err)
		return
	}

	// Uint32传入一个切片返回一个uint32类型长度
	relay := binary.BigEndian.Uint32(buf[:4])
	// 根据relay读取消息内容
	n,err := conn.Read(buf[:relay])
	if n != int(relay) || err != nil{
		// fmt.Println("conn.Read file err = ",err)
		return
	}

	// 这里序列化buf[:relay]并存入msg中一定要加&符号，地址否则不会返回给函数返回值
	err = json.Unmarshal(buf[:relay],&msg)
	if err != nil{
		fmt.Println("json.Unmarshal err = ",err)
		return
	}
	return
}

// 编写一个writePkg函数用来发送请求
func writePkg(conn net.Conn,data []byte)(err error){
	// 先把data转为uint32再转换并放入切片中把切片发送给对方
	var relay uint32
	relay = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4],relay)
	n,err := conn.Write(buf[0:4])
	if err != nil || n != 4{
		fmt.Println("conn.Write err = ",err)
		return
	}

	// 发送data本身
	n,err = conn.Write(data)
	if n != int(relay) || err != nil{
		fmt.Println("writePkg conn.Write err",err)
		return
	}
	return
}
