package message

// 声明两个常量表示LoginMsg和LoginResMsg类型
const(
	LoginMsgType = "LoginMsg" // 登录
	LoginResMsgType = "LoginResMsg" // 返回登录
	RegisterMsgType = " RegisterMsg" // 注册
	RegisterResMsgType = " RegisterResMsg" // 返回注册
	NotifyUserStatusMsgType = "NotifyUserStatusMsg" // 推送用户在线信息
	SmsMsgType = "SmsMsg" // client群发消息类型
	PrivateMsgType = "PrivateMsg" // 私聊消息
)

// 定义几个用户状态的常量
const(
	UserOnlien = iota
	UserOffline
	UserBusyStatus
)

// Message和LoginMsg用户存放消息及给消息打tag
type Message struct{
	Type string `json:"type"` // 消息类型
	Data string  `json:"data"` // 消息内容
}

type LoginMsg struct{
	UserId int `json:"userId"` // 用户id
	UserPwd string `json:"userPwd"` // 用户密码
	UserName string `json:"userName"` // 用户名称
}

// LoginResMsg用来表示返回消息发送情况
type LoginResMsg struct{
	Code int `json:"code"` // 状态码200表示成功，500表示未注册
	UsersId []int `json:"users"` // 保存用户Id的slice
	Error string `json:"error"` // 错误信息
}

type RegisterMsg struct{
	User User `json:"user"` // User 类型来源于message.user
}

type RegisterResMsg struct{
	Code int `json:"code"` // 状态码200表示成功，500表示已被注册
	Error string `json:"error"` // 错误信息
}

type NotifyUserStatusMsg struct{
	UserId int `json:"userId"` // 用户id
	Status int `json:"status"` // 用户状态
}

// 增加一个smsProcess群发的struct
type SmsMsg struct{
	Content string `json:"content"` // 消息内容
	User // 匿名结构体来自于user.User
}

type PrivateMsg struct{
	Content string `json:"content"` // 
	MyUserId int `json:"myUserId"`
	User
}