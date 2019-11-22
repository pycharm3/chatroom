package processall

import (
	"fmt"
	"gotcp/chatroom/common/message"
)

// 定义一个客户端要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User,10)

// 在客户端显示当前在线的用户
func outputOnlineUsers(){
	fmt.Println("当前在线用户列表: ")
	// 遍历一把onlineUsers
	for id,_ := range onlineUsers{
		fmt.Println("用户Id:",id)
	}
}

// 编写一个方法处理返回的NotifyUserStatusMsg
func updataUserStatus(notifyUserStatusMsg *message.NotifyUserStatusMsg){
	user,ok := onlineUsers[notifyUserStatusMsg.UserId]
	if !ok{
		user = &message.User{
			UserId : notifyUserStatusMsg.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMsg.Status
	onlineUsers[notifyUserStatusMsg.UserId] = user
}