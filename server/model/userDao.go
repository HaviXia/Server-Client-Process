package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

/*
	在服务端启动后，初始化 UserDao 实例
	把它做成全局的变量，在需要和redis 交互时，直接使用即可
*/
var (
	MyUserDao *UserDao
)

/*
	userDao 提供了哪些方法？

	type User struct{
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
	}
*/
//定义 UserDao 结构体   ----->   !!!完成对  User 结构体的各种操作!!!
/*
	根据结构图，发现，UserDao 要操作 User 时，需要从 链接池中 取出 conn ， 然后通过这个连接 操作 redis 实现对 User 的操作
*/

type UserDao struct {
	//将链接池作为一个字段给 UserDao ，需要的话自己取
	pool *redis.Pool //
}

/*
	将 pool 声明成 redis.Pool 只使用了 pool.Get() 得到一个连接

	使用工厂模式，获取到 UserDao 的实例
*/
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

/*
	UserDao 提供的方法：

			登陆  --  根据用户的 UserId  返回一个 实例 + error
*/
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) { // 想一下，这个为什么是 redis.Conn 而不是 net.Conn ???
	// 通过给定的 Id 到 redis 中查询用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		// 返回的是什么错误？
		if err == redis.ErrNil { // 就是说没有查找到这个 userId 对应的 value
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	// 得到的 res 是一个序列化之后的内容 "{/"/..../"}"
	// 需要将 res 进行反序列化，得到 User 实例 ，这样才能得到对应的 Id 和 Pwd
	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json Unmarshal err:", err)
		return
	}
	return
}

/*
	第二个功能，完成登陆校验,对用户的 UserId  UserPwd 验证
	用户的 Id 和 Pwd 都正确，则返回一个 User 实例

	如果 UserId 或者 UserPwd 有错，则返回对应的错误信息
*/
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	//从 userDao链接池 中取出一个连接
	conn := this.pool.Get() // 取到连接

	// 连接关闭
	defer conn.Close()

	user, err = this.getUserById(conn, userId) //返回值已经定义名称，使用 = 而不是 :=
	if err != nil {
		//可能是用户不存在，也可能是 Unmarshak
		return
	}

	//没有错误，用户的 id 是存在的，并且获取到了
	//获取到的 user 是反序列化之后的 user ，可以获取到 userPwd
	if user.UserPwd != userPwd {
		//密码不匹配，不正确
		err = ERROR_USER_PWD
		return
	}
	return
}
