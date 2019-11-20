// 此文件用来操作user结构体
package model

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"gotcp/chatroom/common/message"
)

// 在服务器启动后就初始化一个UserDao实例
// 把它做成全局变量方便使用
// 声明一个UserDao结构体，实例化一个pool连接池
type rootUserDao struct{
	pool *redis.Pool
}

// 实例化一个 rootUserDao 指针类型
var (
	MyUserDao *rootUserDao
)

// 使用工厂模式创建UserDao实例
// 工厂模式主要是为了程序的隔离性，在资源和使用者之间提供一个中介服务
// 通过工厂模式函数返回一个 pool实例再赋值给 MyUserDao 进行使用
func NewUserDao(pool *redis.Pool)(userDao *rootUserDao){
	userDao = &rootUserDao{
		pool:pool,
	}
	// 返回userDao
	return
}

// 根据id返回User
// 第一层对信息校验，校验是否能查到对应的ID
func (this *rootUserDao)getUserById(conn redis.Conn,id int)(user *User,err error){
	res,err := redis.String(conn.Do("hget","users",id))
	// 如果 err!= nil 则说明改id未被注册
	if err != nil{
		if err == redis.ErrNil{
			err = ERROR_USER_EXISTENCE
		}
		return
	}

	user = &User{}

	err = json.Unmarshal([]byte(res),user)
	if err != nil{
		fmt.Println("json.Unmarshal([]byte(res),user) err=",err)
	}
	return
}

// 根据返回的User进行校验
// 如果id和pwd都正确返回User实例
// 如果id和pwd有误返回对应错误信息
func (this *rootUserDao)Login(userId int,userPwd string)(user *User,err error){
	// 从rootUserDao中取出一根连接
	conn := this.pool.Get()
	defer conn.Close()
	user,err = this.getUserById(conn,userId)
	if err != nil{
		// fmt.Println("this.getUserById(conn,UserId) err",err)
		// return
		// 进行信息校验
		if err == ERROR_USER_EXISTENCE{
			return		
		}else{
			fmt.Println("this.getUserById err",err)
		}
	}
	// 第二层校验校验查出来的用户密码是否正确
	if userPwd != user.UserPwd{
		err = ERROR_USER_PWD
		return
	}
	return
}

func (this *rootUserDao)Register(user *message.User)(err error){
	conn := this.pool.Get()
	defer conn.Close()
	_,err = this.getUserById(conn,user.UserId)
	// 如果getUserById根据用户传入的Id查到了则err返回空用户已注册，查不到返回err说明用户未注册
	if err == nil{
		err = ERROR_USER_EXISTS
		return
	}
	// 先对用户注册信息进行验证查看是否已存在如果未存在则可以继续注册
	data,err := json.Marshal(user)
	if err != nil{
		fmt.Println("json.Marshal err = ",err)
		return
	}

	_,err = conn.Do("hset","users",user.UserId,string(data))
	if err != nil{
		fmt.Println("用户信息入库出错：",err)
		return
	}
	return
}