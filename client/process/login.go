package process

// 客户端登陆成功，跳转到 cilent 洁面

//func login(userId int, userPwd string) (err error) { // 返回一个 err 信息
//	//fmt.Println("userID = %d , userPwd = %s \n", userID, userPwd)
//	//return nil
//
//	//1。连接到服务器
//	conn, err := net.Dial("tcp", "localhost:8889")
//	if err != nil {
//		fmt.Println("客户端 net.Dial() err,", err)
//		return
//	}
//
//	//延时关闭
//	defer conn.Close()
//
//	// 2。准备通过 conn 发送给消息
//	var mes message.Message
//	mes.Type = message.LoginMesType
//
//	//3。创建 LoginMes
//	var loginMes message.LoginMes
//	loginMes.UserId = userID
//	loginMes.UserPwd = userPwd
//
//	//4。将 loginMes 序列化
//	marshal, err := json.Marshal(loginMes) // marshal是一个切片 []byte
//	if err != nil {
//		panic(err)
//		fmt.Println("loginMes 序列化错误")
//		return
//	}
//
//	//赋值,把 []byte 转换成 string
//	mes.Data = string(marshal)
//
//	//将 mes 进行序列化
//	mesData, err := json.Marshal(mes)
//	if err != nil {
//		panic(err)
//		fmt.Println("mes序列化错误")
//		return
//	}
//
//	//mesData 就是我们要发送的 序列化之后的消息
//	// 发送要先发送数据的长度，再发送数据
//
//	//先获取到 mesData 的长度，将长度转换成 表示长度的[]byte切片
//	//conn.Write(len(data)) 报错
//
//	/*
//		怎么把长度转换成 []byte
//
//		binary包 实现了简单的数字与字节序列的转换以及变长值的编解码。
//
//		type ByteOrder interface{
//			...
//			PutUint32([]byte,uint32)
//			PutUint64([]byte,uint64)
//		}
//	*/
//
//	var pkgLen uint32
//	pkgLen = uint32(len(mesData))
//
//	var buf [4]byte
//
//	//把长度转换为[]byte
//	binary.BigEndian.PutUint32(buf[0:4], pkgLen) //
//	n, err := conn.Write(buf[:4])
//
//	if err != nil {
//		panic(err)
//		fmt.Println("conn.Write 失败！", err)
//		return
//	}
//	fmt.Println("长度为:", n)
//	fmt.Printf("发送的内容为:%v", string(mesData))
//	fmt.Println("客户端发送的长度为,", len(mesData))
//
//	//客户端发送消息本身给服务端
//	_, err = conn.Write(mesData)
//	if err != nil {
//		fmt.Println("conn.Write(mesData) err", err)
//		return
//	}
//
//	//处理服务器端返回的消息,read
//	mes, err = utils.readPkg(conn)
//	if err != nil {
//		fmt.Println("client login readPkg() err", err)
//		return
//	}
//	//反序列化 mes 的 Data,编程 LoginResMes
//	var loginResMes message.LoginResMes
//	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
//	if err != nil {
//		fmt.Println("client login loginResMes Unmarshal err~~~", err)
//		return
//	}
//	if loginResMes.Code == 200 {
//		fmt.Println("用户登录成功")
//	} else if loginResMes.Code == 500 {
//		fmt.Println(loginResMes.Err)
//		fmt.Println("用户不存在请注册后登陆")
//	}
//	return
//}
