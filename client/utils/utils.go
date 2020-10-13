package utils

import (
	"DailyGolang/sxt17_socket/tcp用户即时通信/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type Transfer struct {
	/*
		将两个相关方法封装到结构体中，实现 封层 MVC 和 OOP

		后面可能考虑：封装？？
	*/
	Conn net.Conn

	// Buf 当成切片用
	Buf [8064]byte // Buf 是传输时使用的缓存
}

// conn 连接  可以通过 this 获取
func (this *Transfer) WritePkg(data []byte) (err error) {
	// 先发送一个 数据长File Watchers度
	// 发[]byte长度的逻辑
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[:4], pkgLen) // 把 []byte 转换成 uint32

	//发送数据长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(pkgLen) failed", err)
		return
	}
	//发送消息本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) failed", err)
		return
	}
	return
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//为 buf 开辟新的空间
	//buf := make([]byte, 1024)
	//fmt.Println("读取客户端传来的数据")
	_, err = this.Conn.Read(this.Buf[:4]) //假设只传来4个数据

	//conn.Read 在 conn 没有被关闭的情况下就会阻塞
	//如果客户端关闭了 conn，就不会被阻塞
	//那么 服务端就会一直error 不断的 read pkg head err
	if err != nil {
		fmt.Println("conn.Read failed ,err:", err) //err EOF
		//errors.New("read pkg head error")
		return
	}
	//根据读到的 buf[:4] 转换成 uint32 输出读取到的字节数
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4]) // 把 []byte 转换成 uint32

	//根据 pkgLen 读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen]) // 这句话的理解为，从 conn 中 Read len(pkgLen) 个字节，并放到 buf 里面
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read丢包了,", err)
		//errors.New("read pkg body error")
		return
	}
	// buf 中存储的就是传来的字符串
	// 把pkgLen反序列化成 Message，再对Message反序列化，得到LoginMesStr
	//var mes message.Message // 这个 mes 是一个 结构体类型
	/*func Unmarshal(data []byte, v interface{}) error {}*/
	err = json.Unmarshal(this.Buf[:pkgLen], &mes) //Unmarshal两个参数，一个是 []byte,一个是 类型
	if err != nil {
		fmt.Println("json.Unmarshal failed", err)
		// errors.New("read pkg tail error")
		return
	}
	return
}
