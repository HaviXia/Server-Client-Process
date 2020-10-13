package process

import (
	"DailyGolang/sxt17_socket/tcp用户即时通信/client/utils"
	"fmt"
	"net"
	"os"
)

/*
	显示登陆成功的界面、保持和服务器的通讯、当读取服务器发送的消息后，就会显示在client的界面
*/

//1. 显示登陆成功后的界面

func ShowMenu() {

	// fmt.Printf("恭喜%s登陆成功\n",userName)
	fmt.Println("1。显示在线用户的列表")
	fmt.Println("2。发送消息")
	fmt.Println("3。消息列表")
	fmt.Println("4。退出系统")
	fmt.Println("请选择输入 1 - 4 :")

	var key int
	fmt.Scanln(&key)
	//fmt.Scanf("%d\n",&key)
	switch key {
	case 1:
		//
		fmt.Println("----显示用户在线列表----")
	case 2:
		fmt.Println("----发送消息----")
	case 3:
		fmt.Println("----查看消息列表----")
	case 4:
		fmt.Println("您已经退出系统")
		os.Exit(0) //
	default:
		fmt.Println("输入选项内容错误，请重新输入")
	}
}

// 处理服务器发送来的数据信息
func serverProcessMes(conn net.Conn) {
	// 创建一个 Transfer 实例,不停的读消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Printf("客户端%s正在不停的等待读取服务端传来的消息", conn.LocalAddr())
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg() ,err ~~~~~", err)
			break
			return
		}

		// 代码没有出错的话
		fmt.Printf("mes = %v\n", mes)

	}
}
