package process

import (
	"DailyGolang/sxt17_socket/tcp用户即时通信/common/message"
	/*
		导入外部的包的时候，不要把外部的包 定义为 package main

		main 包无法被其他包 import
	*/
	"DailyGolang/sxt17_socket/tcp用户即时通信/server/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

// 同理 定义一个 UserProcess 结构体，并且绑定方法
type UserProcess struct {
	/*
		Conn 连接  这样以后使用到 conn 连接，直接 this.conn 就直接可以
	*/
	Conn net.Conn
}

// 创建 UserProcess 中包含了 Conn 连接，而 userProcess 中的连接，一定是从  server/main/processor.go(主控) 中传过来的 Conn
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//1 先从 mes 中取出 mes.Data ，并反序列化成 LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal failed err = ", err)
		return
	}

	//声明一个 resMes
	var resMes message.Message
	resMes.Type = message.LoginResType

	//声明一个 LoginResMes,完成赋值
	var loginResMes message.LoginResMes

	// ID = 100 , Pwd = 123456 登陆成功
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		//成功
		loginResMes.Code = 200
		//fmt.Println("登陆成功")
	} else {
		loginResMes.Code = 500
		loginResMes.Err = errors.New("用户登录失败，账号或密码错误")
	}
	//将 loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("loginResMes序列化错误！", err)
		return err
	}

	// 序列化之后，传给resMes
	resMes.Data = string(data)
	// 序列化 resMes 发送
	marData, err := json.Marshal(resMes) // marData 是 []byte，传入到writePkg的参数要是  []byte
	if err != nil {
		fmt.Println("resMes序列化错误！", err)
		return err
	}

	//发送，防止丢包，封装到 writePkg(net.Conn, []byte)

	/*
		改进的bug：因为使用了分层的模式，需要先创建一个 Transfer 实例，然后 transfer.writePkg()
	*/
	// 使用分层MVC

	transfer := &utils.Transfer{
		Conn: this.Conn,
	}
	// 连接的方式
	err = transfer.WritePkg(marData) // 找不到 conn ，需要修改

	return err
}
