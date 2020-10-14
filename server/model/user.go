package model

/*
	定义用户的结构体
*/
type User struct {
	/*
		确定字段信息,之后这几个字段会进行序列化的操作，添加到redis的字段名都是小写的，所以要使用反射

		为了序列化和反序列化成功

		用户信息的 json 字符串的 key 和 结构体的字段对应的 tag 必须名字一致
	*/

	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}
