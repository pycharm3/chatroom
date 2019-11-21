package message

// 声明两个常量表示LoginMsg和LoginResMsg类型
const(
	LoginMsgType = "LoginMsg" // 登录
	LoginResMsgType = "LoginResMsg" // 返回登录
	RegisterMsgType = " RegisterMsg" // 注册
	RegisterResMsgType = " RegisterResMsg" // 返回注册
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