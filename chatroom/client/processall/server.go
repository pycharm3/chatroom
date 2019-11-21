package processall

import (
	"fmt"
	"os"
	"net"
	"gotcp/chatroom/client/utils"
)
// 发送消息 信息列表 退出系统
func Showmenu(){
	for{
		fmt.Println("-----恭喜xxx登录成功-----")
		fmt.Println("-----1 查看在线用户列表-----")
		fmt.Println("-----2 发送消息-----")
		fmt.Println("-----3 信息列表-----")
		fmt.Println("-----4 退出系统-----")
		fmt.Println("请选择1-4")
	
		key := 1
		fmt.Scanf("%d\n",&key)
	
		switch key{
			case 1:
				fmt.Println("查看在线用户列表")
			case 2:
				fmt.Println("发送消息")
			case 3:
				fmt.Println("信息列表")
			case 4:
				fmt.Println("退出系统")
				os.Exit(0)
			default:
				fmt.Println("输入有误请重新输入...")
				fmt.Println("key =",key)
		}	
	}
}

// 和服务器保持通讯
func Stayconnected(conn net.Conn){
	tf := &utils.Transfer{
		Conn : conn,
	}
	for {
		fmt.Println("客户端正在等待读取server传来的消息...")
		msg,err := tf.ReadMsg()
		if err != nil{
			fmt.Println("tf.ReadMsg err = ",err)
		}
		fmt.Println("msg = ",msg)
	} 
}
