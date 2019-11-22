package main

import (
	"io"
	"fmt"
	"net"
	"gotcp/chatroom/common/message"
	"gotcp/chatroom/server/process2"
	"gotcp/chatroom/server/utils"
)

type Processall struct{
	Conn net.Conn `json:"conn"`
}

// 编写一个ServerProcessMsg函数判断消息请求类型进行相应处理
func (this *Processall)serverProcessMsg(msg *message.Message)(err error){
	fmt.Println("看看能否接收到群发消息: ",msg)
	switch msg.Type{
		case message.LoginMsgType:
			up := &process2.Userprocess{
				Conn : this.Conn,
			}
			err = up.ServerProcessLogin(msg)
		case message.RegisterMsgType:
			// 处理注册请求
			up2 := &process2.Userprocess{
				Conn : this.Conn,
			}
			err = up2.ServerProcessRegister(msg)
		default :
			fmt.Println("消息类型不存在")
	}
	return 
}

func (this *Processall)processalltwo(){
	for {
		tf := &utils.Transfer{
			Conn : this.Conn,
		}
		
		msg,err := tf.ReadMsg()
		if err != nil{
			// fmt.Println("readMsg(conn) err = ",err)
			if err == io.EOF{
				fmt.Println("客户端退出服务端也退出...")
				return
			}else{
				fmt.Println("tf.ReadMsg() err = ",err)
				return
			}
		}

		fmt.Println("msg反序列化后 = ",msg)

		err = this.serverProcessMsg(&msg)
		if err != nil{
			fmt.Println("ServerProcessMsg(conn,msg) err:",err)
		}
	}
}