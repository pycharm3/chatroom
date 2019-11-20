// 处理和用户相关的请求，登录，注册，注销，用户列表管理
package process2

import (
	"fmt"
	"net"
	"gotcp/chatroom/common/message"
	"gotcp/chatroom/server/utils"
	"encoding/json"
	"gotcp/chatroom/server/model"
)

type Userprocess struct{
	Conn net.Conn
}

// 编写一个serverProcessLogin处理登录请求
func (this *Userprocess)ServerProcessLogin(msg *message.Message)(err error){
	// 核心代码 先从msg中取出msg.Data并直接反序列化为loginMsg
	// 用作对消息进行判断进而返回client消息
	var loginMsg message.LoginMsg
	err = json.Unmarshal([]byte(msg.Data),&loginMsg)
	if err != nil{
		fmt.Println("serverProcessLogin json.Marshal err:",err)
		return
	}

	// 声明一个Message标识消息类型
	var resMsg message.Message
	resMsg.Type = message.LoginResMsgType

	// 声明一个LoginResMsg完成赋值
	var loginResMsg message.LoginResMsg

	// 使用model.MyUserDao去redis验证
	user,err := model.MyUserDao.Login(loginMsg.UserId,loginMsg.UserPwd)
	if err != nil{
		if err == model.ERROR_USER_EXISTENCE{
			// 用户不存在返回500
			loginResMsg.Code = 500
			loginResMsg.Error = err.Error()
		}else if err == model.ERROR_USER_PWD{
			// 密码不正确返回300
			loginResMsg.Code = 300
			loginResMsg.Error = err.Error()
		}else{
			loginResMsg.Code = 500
			loginResMsg.Error = "内部信息错误"
		}
	}else{
		loginResMsg.Code = 200
		fmt.Println(user,"登录成功")
	}


	// if loginMsg.UserId == 100 && loginMsg.UserPwd == "123456"{
	// 	loginResMsg.Code = 200
	// }else{
	// 	loginResMsg.Code = 500
	// 	fmt.Println("user notfund fpx.no1")
	// }

	data,err := json.Marshal(loginResMsg)
	if err != nil{
		fmt.Println("serverProcessLogin json.Marshal(loginResMsg) err:",err)
		return
	}

	resMsg.Data = string(data)
	data,err = json.Marshal(resMsg)
	if err != nil{
		fmt.Println("serverProcessLogin json.Marshal(resMsg) err:",err)
		return
	}

	// 消息序列化完毕准备发送
	tf := &utils.Transfer{
		Conn : this.Conn, // 调用本方法(this *Userprocess) this绑定struct里的Conn
	}

	err = tf.WritePkg(data)
	return
}

// 处理注册请求
func (this *Userprocess)ServerProcessRegister(msg *message.Message)(err error){
	var registermsg message.RegisterMsg
	err = json.Unmarshal([]byte(msg.Data),&registermsg)
	if err != nil{
		fmt.Println("serverProcessLogin json.Marshal err:",err)
		return
	}

	var resMsg message.Message
	resMsg.Type = message.RegisterResMsgType
	var registerResMsg message.RegisterResMsg

	err = model.MyUserDao.Register(&registermsg.User)
	if err != nil{
		if err == model.ERROR_USER_EXISTS{
			registerResMsg.Code = 505
			registerResMsg.Error = model.ERROR_USER_EXISTS.Error()
		}else{
			registerResMsg.Code = 200
		}
	}

	data,err := json.Marshal(registerResMsg)
	if err != nil{
		fmt.Println("serverProcessLogin json.Marshal(registerResMsg) err:",err)
		return
	}

	resMsg.Data = string(data)
	data,err = json.Marshal(resMsg)
	if err != nil{
		fmt.Println("serverProcessLogin json.Marshal(resMsg) err:",err)
		return
	}

	// 消息序列化完毕准备发送
	tf := &utils.Transfer{
		Conn : this.Conn, // 调用本方法(this *Userprocess) this绑定struct里的Conn
	}
	err = tf.WritePkg(data)
	return
}