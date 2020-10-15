package message

const (
	LoginMesType       = "LoginMes"
	LoginResMesType    = "LoginResMes"
	RegisterMesType    = "RegisterMes"
	RegisterResMesType = "RegisterResMes"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// 定义消息

/*
	登陆的消息（userID、userPwd）
	登陆结果的消息（err error）
	注册的消息 ()
*/

type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	Code int `json:"code"` //错误的代码
	/*
		Code = 500 表示用户还未注册，不存在
		Code = 403 表示用户密码不正确
		Code = 505 未知错误/服务器内部错误
		Code = 200 表示登陆成功
	*/
	Error string `json:"error"` //登陆时返回的错误信息
}

//注册对应的结构体
type RegisterMes struct {
	User User `json:"user"` // 传入的就是 User 结构体
}

// 描述注册消息类型

type RegisterResMes struct {
	Code int `json:"code"` //错误的代码
	/*
		Code = 400 表示用户Id被占用
		Code = 200 表示登陆成功
	*/
	Error string `json:"error"` //登陆时返回的错误信息
}
