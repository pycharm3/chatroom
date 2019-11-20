// 2019-11-7 下午二点 系统index首页
package main

import (
	"fmt"
	"os"
)

var (
	userid int
	userpwd string
)

func main(){
	var loop = true
	var work int

	// 这里真是没想到，loop=true则为死循环，loop=false退出死循环
	for loop{
		fmt.Println("----------欢迎登录多人聊天系统----------")
		fmt.Println("			1 登录聊天系统")
		fmt.Println("			2 注册用户")
		fmt.Println("			3 退出系统")
		fmt.Println("请输入1-3选择服务")
		fmt.Scanln(&work)
		switch work{
		case 1:
			fmt.Println("---登录聊天系统---")
			loop = false
		case 2:
			fmt.Println("---注册用户---")
			loop = false
		case 3:
			fmt.Println("---退出系统---")
			// os.Exit可以用于退出当前程序
			os.Exit(0)
		default :
			fmt.Println("请输入正确选项1-3")
		}
	}

	if work == 1{
		fmt.Println("请输入用户id:")
		fmt.Scanf("%d\n",&userid)
		fmt.Println("请输入用户密码:")
		fmt.Scanf("%s",&userpwd)

		err := login(userid,userpwd)
		if err != nil{
			fmt.Println("login(userid,userpwd) err:",err)
			return
		}
		
		// if err != nil{
		// 	fmt.Println("登录失败")
		// }else{
		// 	fmt.Println("登录成功!")
		// }

	}else if work != 1{
		fmt.Println("其他选项...")
	}
}