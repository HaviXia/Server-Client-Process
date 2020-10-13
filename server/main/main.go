package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("系统实现了结构分层之后的DEMO～")
	fmt.Println("服务器在 8889 端口监听.......")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("服务器监听8889端口错误,", err)
		return
	}
	defer listen.Close()

	// 监听成功 ，等待客户端连接服务器
	for {
		fmt.Println("等待客户端连接服务器")
		conn, err := listen.Accept() // conn 就是服务器和客户端之间互通的 连接
		if err != nil {
			fmt.Println("Listen accept() 出错,", err)
			return
		}

		// 拿到了连接，启动一个协程 和 客户端保持通讯
		go process(conn)
	}
}

// 处理和 客户端的通讯 , 传入的参数 要是一个连接
func process(conn net.Conn) {

	defer conn.Close() //延时关闭客户端的连接

	// 由于把 ReadPkg() 写在了 processor.go 中，所以这里需要调用 processor.go
	processor := &Processor{
		Conn: conn, // 这个连接会一层一层的往下传
	}
	err := processor.processMain()
	if err != nil {
		fmt.Println("客户端和服务器端通讯的协程出现了问题", err)
		return
	}
	//buf := make([]byte, 1024*4)

	//// 读取客户端发送过来的信息数据
	//for {
	//
	//	// 第二步，这里我们将读取数据包，直接封装成一个函数 readPkg(),返回Message Err
	//
	//	mes, err := readPkg(conn)
	//	if err != nil {
	//
	//		// 防止客户端关闭了conn，服务端一直 conn head error
	//		if err == io.EOF {
	//			fmt.Println("客户端关闭了连接，服务端也需要退出连接")
	//			return
	//		} else {
	//			fmt.Println("readPkg失败，", err)
	//			return
	//		}
	//		return
	//	}
	//	fmt.Println("mes=", mes)
	//	/*	 定义一个 buffer
	//	fmt.Println("等待读取客户端传来的数据.....")
	//	n, err := conn.Read(buf[:4])
	//	if n != 4 || err != nil {
	//		fmt.Println("服务端 conn.Read err", err)
	//		return
	//	}
	//	fmt.Println("读取到的 buf=", buf[:4])*/
	//	err = serverProcessMes(conn, &mes)
	//	if err != nil {
	//		return
	//	}
	//}

}

//第二步，定义的一个新的函数
//func readPkg(conn net.Conn) (mes message.Message, err error) {
//	//为 buf 开辟新的空间
//	buf := make([]byte, 1024)
//	fmt.Println("读取客户端传来的数据")
//	_, err = conn.Read(buf[:4]) //假设只传来4个数据
//
//	//conn.Read 在 conn 没有被关闭的情况下就会阻塞
//	//如果客户端关闭了 conn，就不会被阻塞
//	//那么 服务端就会一直error 不断的 read pkg head err
//	if err != nil {
//		fmt.Println("conn.Read failed ,err:", err) //err EOF
//		//errors.New("read pkg head error")
//		return
//	}
//	//根据读到的 buf[:4] 转换成 uint32 输出读取到的字节数
//	var pkgLen uint32
//	pkgLen = binary.BigEndian.Uint32(buf[:4]) // 把 []byte 转换成 uint32
//
//	//根据 pkgLen 读取消息内容
//	n, err := conn.Read(buf[:pkgLen]) // 这句话的理解为，从 conn 中 Read len(pkgLen) 个字节，并放到 buf 里面
//	if n != int(pkgLen) || err != nil {
//		fmt.Println("conn.Read丢包了,", err)
//		errors.New("read pkg body error")
//		return
//	}
//	// buf 中存储的就是传来的字符串
//	// 把pkgLen反序列化成 Message，再对Message反序列化，得到LoginMesStr
//	//var mes message.Message // 这个 mes 是一个 结构体类型
//	/*func Unmarshal(data []byte, v interface{}) error {}*/
//	err = json.Unmarshal(buf[:pkgLen], &mes) //Unmarshal两个参数，一个是 []byte,一个是 类型
//	if err != nil {
//		fmt.Println("json.Unmarshal failed", err)
//		errors.New("read pkg tail error")
//		return
//	}
//	return
//}

// 根据客户端发送的消息种类不同，决定调用哪个函数处理
//func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
//
//	switch mes.Type {
//	case message.LoginMesType:
//		// 处理登陆
//		err = serverProcessLogin(conn, mes)
//	case message.RegisterType:
//		//
//
//	default:
//		fmt.Println("消息类型不存在，无法处理.....")
//	}
//	return
//}

//func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
//	//1 先从 mes 中取出 mes.Data ，并反序列化成 LoginMes
//	var loginMes message.LoginMes
//	err = json.Unmarshal([]byte(mes.Data), &loginMes)
//	if err != nil {
//		fmt.Println("json.Unmarshal failed err = ", err)
//		return
//	}
//
//	//声明一个 resMes
//	var resMes message.Message
//	resMes.Type = message.LoginResType
//
//	//声明一个 LoginResMes,完成赋值
//	var loginResMes message.LoginResMes
//
//	// ID = 100 , Pwd = 123456 登陆成功
//	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
//		//成功
//		loginResMes.Code = 200
//		fmt.Println("登陆成功")
//	} else {
//		loginResMes.Code = 500
//		loginResMes.Err = errors.New("用户登录失败，账号密码错误")
//
//		//将 loginResMes 序列化
//		data, err := json.Marshal(loginResMes)
//		if err != nil {
//			fmt.Println("loginResMes序列化错误！", err)
//
//		}
//
//		// 序列化之后，传给resMes
//		resMes.Data = string(data)
//		// 序列化 resMes 发送
//		marData, err := json.Marshal(resMes) // marData 是 []byte，传入到writePkg的参数要是  []byte
//		if err != nil {
//			fmt.Println("resMes序列化错误！", err)
//
//		}
//
//		//发送，防止丢包，封装到 writePkg(net.Conn, []byte)
//		err = writePkg(conn, marData)
//		if err != nil {
//			panic(err)
//
//		}
//	}
//	return
//}
//func writePkg(conn net.Conn, data []byte) (err error) {
//	// 先发送一个 数据长File Watchers度
//	// 发[]byte长度的逻辑
//	var pkgLen uint32
//	pkgLen = uint32(len(data))
//	var buf [4]byte
//	binary.BigEndian.PutUint32(buf[:4], pkgLen) // 把 []byte 转换成 uint32
//
//	//发送数据长度
//	n, err := conn.Write(buf[:4])
//	if n != 4 || err != nil {
//		fmt.Println("conn.Write(pkgLen) failed", err)
//		return
//	}
//	//发送消息本身
//	n, err = conn.Write(data)
//	if n != int(pkgLen) || err != nil {
//		fmt.Println("conn.Write(data) failed", err)
//		return
//	}
//	return
//}
