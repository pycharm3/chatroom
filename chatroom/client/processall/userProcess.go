package processall

import (
	"fmt"
	"net"
	"gotcp/chatroom/common/message"
	"gotcp/chatroom/client/utils"
	"encoding/json"
	"os"
)

type Lucky struct{

}

func (this *Lucky)Login(userid int,userpwd string)(err error){

	// 下一步制定协议
	conn,err := net.Dial("tcp","0.0.0.0:8889")
	defer conn.Close()

	if err != nil{
		fmt.Println("net.Dial err = ",err)
		return
	}

	// 实例化一个utils下的Transfer
	tf := &utils.Transfer{
		Conn : conn,
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
	fmt.Println("data序列化后 = ",string(data))
	// 调用WritePkg发送消息长度及消息本身
	err = tf.WritePkg(data)
	if err != nil{
		fmt.Println("conn.Write err = ",err)
	}
	fmt.Println("client message send ok!")
	// 休眠20s
	// time.Sleep(20 * time.Second)
	// fmt.Println("休眠了20s...")

	// 这里还需要处理服务端返回的消息
	msg,err = tf.ReadMsg()
	if err != nil{
		fmt.Println("readMsg(conn) err :",err)
	}

	// 将msg.Date反序列化
	var loginResmsg message.LoginResMsg

	err = json.Unmarshal([]byte(msg.Data),&loginResmsg)

	if loginResmsg.Code == 200{
		for _,v := range loginResmsg.UsersId{
			// 如果v等于自己登录的Id则使用continue退出当前循环
			if v == userid{
				continue
			}
			fmt.Println("登录成功...当前在线用户Id为：",v)
		}
		/* 登录成功调用Stayconnected()和服务器保持连接，有信息传入
		client则Showmenu()能看到 */
		// 这里开一个goroutine去单独执行
		go Stayconnected(conn)
		// 循环显示菜单
		Showmenu()
	}else{
		fmt.Println(loginResmsg.Error)
	}
	return
}

// 注册方法
func (this *Lucky)Register(userid int,username string,userpwd string)(err error){
	conn,err := net.Dial("tcp","0.0.0.0:8889")
	defer conn.Close()

	if err != nil{
		fmt.Println("net.Dial err = ",err)
		return
	}
	// 实例化一个消息struct并将消息类型放入并将用户信息序列化后放入
	var msg message.Message
	msg.Type = message.RegisterMsgType

	// 实例化一个用户信息struct并将用户信息放入实例中并序列化
	var registermsg message.RegisterMsg 
	registermsg.User.UserId = userid
	registermsg.User.UserName = username
	registermsg.User.UserPwd = userpwd

	data,err := json.Marshal(registermsg)
	if err != nil{
		fmt.Println("json.Marshal err = ",err)
		return
	}
	
	msg.Data = string(data)

	// 将msg.struct序列化
	data,err = json.Marshal(msg)
	if err != nil{
		fmt.Println("json.Marshal err = ",err)
		return
	}
	fmt.Println("data序列化后 = ",string(data))

	// 实例化一个utils下的Transfer
	tf := &utils.Transfer{
		Conn : conn,
	}
	err = tf.WritePkg(data)

	if err != nil{
		fmt.Println("conn.Write err = ",err)
		return
	}
	fmt.Println("client message send ok!")

	// 接收server响应
	msg,err = tf.ReadMsg()
	if err != nil{
		fmt.Println("readMsg(conn) err :",err)
		return
	}
	// 将msg.Date反序列化
	var registerResMsg message.RegisterResMsg
	err = json.Unmarshal([]byte(msg.Data),&registerResMsg)
	if registerResMsg.Code == 200{
		fmt.Println("注册成功...")
	}else{
		fmt.Println(registerResMsg.Error)
	}
	os.Exit(0)
	return
}