package main

import (
	"DailyGolang/sxt17_socket/tcp用户即时通信/common/message"
	usrProcess "DailyGolang/sxt17_socket/tcp用户即时通信/server/process"
	"DailyGolang/sxt17_socket/tcp用户即时通信/server/utils"
	"io"

	"fmt"
	"net"
)

//	创建 processor 的结构体
type Processor struct {
	/*
		需要连接
	*/
	Conn net.Conn
}

// 判断 登陆的类型
func (this *Processor) ServerProcessMes(mes *message.Message) (err error) {

	switch mes.Type {
	case message.LoginMesType:
		// 处理登陆
		/*
			分层之后，需要先创建一个 userProcess 实例，调用其 ServerProcessLogin 的方法
		*/
		usrProcess := &usrProcess.UserProcess{Conn: this.Conn}
		err = usrProcess.ServerProcessLogin(mes)
	case message.RegisterType:

		// 处理登陆的逻辑

	default:
		fmt.Println("消息类型不存在，无法处理.....")
	}
	return
}

func (this *Processor) processMain() (err error) {
	//defer this.Conn.Close()
	// 读取客户端发送过来的信息数据
	for {
		// 创建引用
		transfer := &utils.Transfer{
			Conn: this.Conn,
		}

		mes, err := transfer.ReadPkg()
		if err != nil {

			// 防止客户端关闭了conn，服务端一直 conn head error
			if err == io.EOF {
				fmt.Println("客户端关闭了连接，服务端也需要退出连接")
				return err
			} else {
				fmt.Println("readPkg失败，", err)
				return err
			}
			return err
		}
		//fmt.Println("mes=", mes)
		/*	 定义一个 buffer
		fmt.Println("等待读取客户端传来的数据.....")
		n, err := conn.Read(buf[:4])
		if n != 4 || err != nil {
			fmt.Println("服务端 conn.Read err", err)
			return
		}
		fmt.Println("读取到的 buf=", buf[:4])*/
		err = this.ServerProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
