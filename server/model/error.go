package model

import "errors"

/*
	根据业务逻辑的需要，自定义一些错误 error
*/

var (
	/*
		用户不存在 / 用户存在 / 密码错误
	*/
	ERROR_USER_NOTEXISTS = errors.New("该用户不存在")
	ERROR_USER_EXISTS    = errors.New("用户已经存在") // 注册时，ID唯一
	ERROR_USER_PWD       = errors.New("密码不正确")
)

type Error struct {
}
