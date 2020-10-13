package message

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
		Code = 500 表示用户还未注册
		Code = 200 表示登陆成功
	*/
	Err error `json:"err"` //登陆时返回的错误信息
}

//注册对应的结构体
type RegisterMes struct {
}

// 描述消息是什么类型的  LoginMes 和 LoginResMes 两种类型

const (
	LoginMesType = "LoginMes"
	LoginResType = "LoginResMes"
	RegisterType = "RegisterMes"
)
