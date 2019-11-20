package utils

import (
	"fmt"
	"net"
	"gotcp/chatroom/common/message"
	"encoding/binary"
	"encoding/json"
)

type Transfer struct{
	Conn net.Conn
	Buf [9600]byte // 传输时使用缓冲
}

func (this *Transfer)ReadMsg()(msg message.Message,err error){
	// buf := make([]byte,9600)
	fmt.Println("准备读取服务端传输数据...")
	// 从conn中读取4个字节放入bug中
	_,err = this.Conn.Read(this.Buf[:4])
	if err != nil{
		// fmt.Println("conn.Read err = ",err)
		return
	}

	// Uint32传入一个切片返回一个uint32类型长度
	relay := binary.BigEndian.Uint32(this.Buf[:4])
	// 根据relay读取消息内容
	n,err := this.Conn.Read(this.Buf[:relay])
	if n != int(relay) || err != nil{
		// fmt.Println("conn.Read file err = ",err)
		return
	}

	// 这里序列化buf[:relay]并存入msg中一定要加&符号，地址否则不会返回给函数返回值
	err = json.Unmarshal(this.Buf[:relay],&msg)
	if err != nil{
		fmt.Println("json.Unmarshal err = ",err)
		return
	}
	return
}

// 编写一个writePkg函数用来发送请求
func (this *Transfer)WritePkg(data []byte)(err error){
	// 先把data转为uint32再转换并放入切片中把切片发送给对方
	var relay uint32
	relay = uint32(len(data))
	// var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[:4],relay)
	n,err := this.Conn.Write(this.Buf[0:4])
	if err != nil || n != 4{
		fmt.Println("conn.Write err = ",err)
		return
	}

	// 发送data本身
	n,err = this.Conn.Write(data)
	if n != int(relay) || err != nil{
		fmt.Println("writePkg conn.Write err",err)
		return
	}
	return
}


