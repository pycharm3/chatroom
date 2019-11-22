package process2

import (
	"net"
	"fmt"
	"encoding/json"
	"gotcp/chatroom/common/message"
	"gotcp/chatroom/client/utils"
)

type SmsProcess struct{

}

func (this *SmsProcess)SendGroupMsg(msg *message.Message){
	var smsMsg message.SmsMsg
	err := json.Unmarshal([]byte(msg.Data),&smsMsg)
	if err != nil{
		fmt.Println("json.UnMarshal([]byte(msg.Data),&smsMsg) err = ",err)
	}
	for id,up := range userMgr.onlineUsers{
		if id == smsMsg.UserId{
			continue
		}
		data,err := json.Marshal(msg)
		if err != nil{
			fmt.Println("json.Marshal(msg) err = ",err)
		}
		this.SendMsgToEachOnlineUser(data,up.Conn)
	}
}

func (this *SmsProcess)SendMsgToEachOnlineUser(data []byte,conn net.Conn){
	tf := &utils.Transfer{
		Conn : conn,
	}
	err := tf.WritePkg(data)
	if err != nil{
		fmt.Println("tf.WritePkg(data) err = ",err)
	}
}