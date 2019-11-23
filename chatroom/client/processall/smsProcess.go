package processall

import (
	"fmt"
	"gotcp/chatroom/common/message"
	"gotcp/chatroom/client/utils"
	"encoding/json"
)

type SmsProcess struct{

}

func (sms *SmsProcess)SendGroupMsg(content string)(err error){
	var msg message.Message
	msg.Type = message.SmsMsgType
	
	var smSmsg message.SmsMsg
	smSmsg.Content = content
	smSmsg.UserId = CurUser.UserId
	smSmsg.UserStatus = CurUser.UserStatus
	// 序列化smSmsg
	data,err := json.Marshal(smSmsg)
	if err != nil{
		fmt.Println("json.Marshal(smSmsg) err = ",err)
		return
	}
	msg.Data = string(data)

	data,err = json.Marshal(msg)
	if err != nil{
		fmt.Println("json.Marshal(msg) err = ",err)
		return
	}

	tf := &utils.Transfer{
		Conn : CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil{
		fmt.Println("tf.WritePkg(data) err = ",err)
		return
	}
	return
}

// 发送私聊信息
func (this *SmsProcess)SendPrivateMsg(content string,userid int)(err error){
	var msg message.Message
	msg.Type = message.PrivateMsgType
	
	var privateMsg message.PrivateMsg
	privateMsg.Content = content
	privateMsg.UserId = CurUser.UserId
	privateMsg.MyUserId = userid
	privateMsg.UserStatus = CurUser.UserStatus
	data,err := json.Marshal(privateMsg)
	if err != nil{
		fmt.Println("json.Marshal(privateMsg) err = ",err)
		return
	}
	msg.Data = string(data)
	data,err = json.Marshal(msg)
	if err != nil{
		fmt.Println("json.Marshal(msg) err = ",err)
		return
	}
	tf := &utils.Transfer{
		Conn : CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil{
		fmt.Println("tf.WritePkg(data) err = ",err)
		return
	}
	return
}