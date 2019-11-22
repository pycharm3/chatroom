package model

import (
	"net"
	"gotcp/chatroom/common/message"
)

// 将CurUser作为一个全局使用
type CurUser struct{
	Conn net.Conn `json:"conn"`
	message.User // 继承
}
