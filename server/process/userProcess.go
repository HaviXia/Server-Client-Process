package process

import (
	"DailyGolang/sxt17_socket/tcp用户即时通信/common/message"
	"DailyGolang/sxt17_socket/tcp用户即时通信/server/model"
	/*
		导入外部的包的时候，不要把外部的包 定义为 package main

		main 包无法被其他包 import
	*/
	"DailyGolang/sxt17_socket/tcp用户即时通信/server/utils"
	"encoding/json"
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
	resMes.Type = message.LoginResMesType

	//声明一个 LoginResMes,完成赋值
	var loginResMes message.LoginResMes

	// 去 redis 验证
	// 在 main 中 实例化了 model.MyUserDao, 使用 model.MyUserDao 去 redis 中进行验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	// 得到的这个 user 非常重要，为了服务器为了拿到用户的各种信息，这个user中包含了id、pwd、name等各种信息

	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			// 可以根据具体的错误信息，返回具体的错误信息
			//loginResMes.Err = err.Error()
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD { //密码不正确
			// 密码不正确定义成 300
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			// 未知信息
			loginResMes.Code = 505
			fmt.Println("服务器内部错误。。。 ")
		}
	} else {
		loginResMes.Code = 200
		fmt.Println("user:登陆成功~\n", user)
	}

	fmt.Printf("user = %v\n", user)
	// ID = 100 , Pwd = 123456 登陆成功
	// 定义了 userDao 现在就要修改原来的验证方式
	//if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	//	//成功
	//	loginResMes.Code = 200
	//	//fmt.Println("登陆成功")
	//} else {
	//	loginResMes.Code = 500
	//	loginResMes.Err = errors.New("用户登录失败，账号或密码错误")
	//}

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

//~~~~~~~~~~~~~~~~~~~~~~~~~~~
func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json Marshal err(),", err)
		return
	}

	//声明一个 resMes ,响应给客户端的消息
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	//声明 registerResMes
	var registerResMes message.RegisterResMes
	// redis 完成注册
	err = model.MyUserDao.Register(&registerMes.User) //两个 user ，要在 Dao 中传入参数是 message.User
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505 //
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 501 //
			registerResMes.Error = "注册发生未知错误"
		}
	} else {
		//fmt.Println("用户注册成功")
		registerResMes.Code = 200

	}

	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("registerResMes序列化错误！", err)
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
	transfer := &utils.Transfer{
		Conn: this.Conn,
	}
	// 连接的方式
	err = transfer.WritePkg(marData)
	return
}
