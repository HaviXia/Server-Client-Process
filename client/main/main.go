package main

import (
	"DailyGolang/sxt17_socket/tcp用户即时通信/client/process"
	"fmt"
	"os"
)

//定义两个变量。一个表示用户的id，一个表示用户的密码

var userId int
var userPwd string
var userName string

// 登陆界面
func main() {
	/*
		欢迎登陆多人聊天系统
	*/

	// 接收用户的菜单的输入
	var key int

	//判断是否继续循环的显示菜单
	//var loop bool = true

	for {
		fmt.Println("----------欢迎登陆多人聊天系统--------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 4 请选择 1-3")

		fmt.Scanf("%d\n", &key) // Scanf 输入必须得保证换行，否则输入之后回车就乱了

		switch key {
		case 1:
			//
			fmt.Println("登陆聊天系统")
			fmt.Println("请输入用户的的id:")
			fmt.Scanln(&userId)
			fmt.Println("请输入用户的的密码:")
			fmt.Scanln(&userPwd)
			//loop = false
			/*
				创建一个 userProcess 的实例
			*/
			up := &process.UserProcess{}
			err := up.Login(userId, userPwd)
			if err != nil {
				//fmt.Println("用户登录失败！", err)
				return
			}

		case 2:
			//
			fmt.Println("注册用户")
			fmt.Println("请输入用户名:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码:")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户的昵称:")
			fmt.Scanf("%s\n", &userName)
			//调用 UserProcess 完成注册
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)

		case 3:
			//
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}
}
