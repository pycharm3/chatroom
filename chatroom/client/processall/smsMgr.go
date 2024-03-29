package processall

import (
	"fmt"
	"gotcp/chatroom/common/message"
	"encoding/json"
)

func outputGroupMsg(msg *message.Message){
	var smsMsg message.SmsMsg
	err := json.Unmarshal([]byte(msg.Data),&smsMsg)
	if err != nil{
		fmt.Println("json.Unmarshal([]byte(msg.Data),&smsMsg) err",err)
		return
	}
	info := fmt.Sprintf("用户Id:\t%v 对大家说:\t%v",smsMsg.UserId,smsMsg.Content)
	fmt.Println(info)
	fmt.Println()
}

func privateMsg(msg *message.Message){
	var privateMsg message.PrivateMsg
	err := json.Unmarshal([]byte(msg.Data),&privateMsg)
	if err != nil{
		fmt.Println("json.Unmarshal([]byte(msg.Data),&privateMsg) err=",err)
		return
	}
	info := fmt.Sprintf("用户Id:\t%v 想对你说:\t%v",privateMsg.UserId,privateMsg.Content)
	fmt.Println(info)
	fmt.Println()
}