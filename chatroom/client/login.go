package main

import (
	"fmt"
	"net"
	"gotcp/chatroom/common/message"
	"encoding/json"
	"encoding/binary"
	_"time"
)

func login(userid int,userpwd string)(err error){
	// 下一步制定协议
	conn,err := net.Dial("tcp","0.0.0.0:8889")
	defer conn.Close()

	if err != nil{
		fmt.Println("net.Dial err = ",err)
		return
	}

	// 实例化一个消息struct并将消息类型放入并将用户信息序列化后放入
	var msg message.Message
	msg.Type = message.LoginMsgType

	// 实例化一个用户信息struct并将用户信息放入实例中并序列化
	var loginmsg message.LoginMsg
	loginmsg.UserId = userid
	loginmsg.UserPwd = userpwd

	// 序列化usermessage
	data,err := json.Marshal(loginmsg)
	if err != nil{
		fmt.Println("json.Marshal err = ",err)
		return
	}

	// 将loginmsg序列化后转为string赋值给msg.Data
	msg.Data = string(data)

	// 将msg.struct序列化
	data,err = json.Marshal(msg)
	if err != nil{
		fmt.Println("json.Marshal err = ",err)
		return
	}

	relay := uint32(len(data))
	var buf [4]byte
	/* Big-Endian就是高位字节排放在内存的低地址端，
	 低位字节排放在内存的高地址端 */
	 // 将uint32类型的长度转化为byte数组并按BigEndian排列于字节数组中
	binary.BigEndian.PutUint32(buf[:4],relay)
	n,err := conn.Write(buf[0:4])
	if err != nil || n != 4{
		fmt.Println("conn.Write err = ",err)
		return
	}
	// 发送消息本身
	n,err = conn.Write(data)
	if err != nil{
		fmt.Println("conn.Write err = ",err)
	}

	fmt.Println("client message send ok!")
	// 休眠20s
	// time.Sleep(20 * time.Second)
	// fmt.Println("休眠了20s...")

	// 这里还需要处理服务端返回的消息
	msg,err = readMsg(conn)
	if err != nil{
		fmt.Println("readMsg(conn) err :",err)
	}

	// 将msg.Date反序列化
	var loginResmsg message.LoginResMsg

	err = json.Unmarshal([]byte(msg.Data),&loginResmsg)

	if loginResmsg.Code == 200{
		fmt.Println("登录成功...")
	}else if loginResmsg.Code == 500{
		fmt.Println("用户不存在...",loginResmsg.Error)
	}
	
	return
}