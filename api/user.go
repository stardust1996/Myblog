package api

import (
	"MyBlog/cache"
	"MyBlog/serializer"
	"MyBlog/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserSignUp(c *gin.Context) {
	var userService service.UserService
	if err := c.ShouldBind(&userService); err == nil {
		res := userService.SignUp()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogin 用户登录接口
func UserSignIn(c *gin.Context) {
	var userService service.UserService
	if err := c.ShouldBind(&userService); err == nil {
		res := userService.SignIn(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	if user := CurrentUser(c);user!=nil {
		res := serializer.BuildUserResponse(*user)
		c.JSON(200, res)
	}else {
		c.JSON(200, gin.H{
			"data":nil,
		})
	}
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	userId := s.Get("user_id"	)
	s.Clear()
	_ = s.Save()
	//redis中删除信息
	cache.RedisClient.Del(userId.(string))
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}
//跳转注册页面
func UserLogin(c *gin.Context) {
	if u := CurrentUser1(c);u != nil{
		//如果已经登录跳转主页
		c.Request.URL.Path="/blogs"
		//GetList(c)
	}else {
		c.HTML(200,"user.html",gin.H{})
	}

}
//获取所有用户
func GetAllUser(c *gin.Context) {
	var userService service.UserService
	re :=  userService.GetAllUser()
	c.JSON(200,re)
}

//删除用户
func DeleteUser(c *gin.Context) {
	var userService service.UserService
	id := c.Param("id")
	re := userService.DeleteUser(id)
	c.JSON(200,re)
}