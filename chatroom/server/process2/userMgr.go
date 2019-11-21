package process2

import (
	"fmt"
)

var (
	// 这里userMgr是pointer类型非常重要，使用userMgr调用UserMgr里的字段
	userMgr *UserMgr 
)

// map value为UserProcess是一个conn连接
type UserMgr struct{
	onlineUsers map[int]*Userprocess
}

// 完成对UserMgr进行初始化
func init(){
	userMgr = &UserMgr{
		// 这里应该是赋值操作，但是map类型的赋值就是make初始化一下
		onlineUsers : make(map[int]*Userprocess,1024),
	}
}

// 完成对onlineUsers添加
func (this *UserMgr)AddOnlineUsers(up *Userprocess){
	this.onlineUsers[up.UserId] = up
}

// 删除
func (this *UserMgr)DelOnlineUsers(userId int){
	delete(this.onlineUsers,userId)
}

// 返回当前所有在线数
func (this *UserMgr)GetAllOnlineUser() map[int]*Userprocess{
	return this.onlineUsers
}

// 根据Id返回对应值
func (this *UserMgr)GetOnlineUserById(userId int)(up *Userprocess,err error){
	// 如何从User中取出一个值，带检测方式
	up,ok := this.onlineUsers[userId]
	if !ok{		// 说明当前用户不在线
		err = fmt.Errorf("用户 %d 不存在",userId)
		return
	}
	return
}