package serializer

import (
	"MyBlog/model"
	"time"
)

// User 用户序列化器
type User struct {
	ID        string     `json:"id"`
	UserName  string     `json:"username"`
	Created   time.Time  `json:"created"`
	Updated   time.Time  `json:"updated"`
}

// BuildUser 序列化用户
func BuildUser( user model.User) User {
	return User{
		ID : user.Id,
		UserName:user.Username,
		Created:user.Created,
		Updated:user.Updated,
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) Response {
	return Response{
		Code: 200,
		Data: BuildUser(user),
	}
}
