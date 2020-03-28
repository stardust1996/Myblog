package server

import (
	"MyBlog/api"
	"MyBlog/middleware"
	"github.com/gin-gonic/gin"
	"os"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	//中间件
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))//设置session
	r.Use(middleware.Cors())//设置跨域验证
	r.Use(middleware.CurrentUser())//获取当前用户

	//载入静态资源
	r.Static("/static","./static")
	r.LoadHTMLGlob("static/views/*")
	//r.GET("/index",api.GetList)//获取博客首页
	//博客相关查询
	r.GET("/blogs",api.GetList)//获取博客列表
	r.GET("/blog/:id",api.GetBlogById)//获取单个博客-
	r.GET("/titles",api.GetBlogTitles)//获取标题
	//获取分类
	r.GET("/classes",api.GetClasses)
	//获取评论
	r.GET("/comments",api.GetComments)
	//跳转用户登录界面
	r.GET("/login",api.UserLogin)
	//用户登录
	r.GET("/user",api.UserSignIn)
	//用户注册
	r.POST("/user",api.UserSignUp)
	//获取用户信息
	r.GET("/user/me",api.UserMe)
	//需要登录才能使用的功能
	auth := r.Group("")
	auth.Use(middleware.AuthRequired())//验证登录的方法
	{
		auth.POST("/comment",api.CreateComments)//添加评论
		auth.DELETE("/logout",api.UserLogout)//登出用户
		//需要管理员权限
		admin := auth.Group("")
		admin.Use(middleware.AdminRequired())//验证管理员身份
		{
			admin.GET("/admin",api.AdminPage)//跳转管理页面
			admin.POST("/class",api.CreateClass)//增加新分类
			admin.POST("/blog",api.CreateBlog)//增加新博客
			admin.DELETE("/blog/:id",api.DeleteBlog)//删除博客
			admin.PUT("/blog",api.UpdateBlog)//更新修改博客
			admin.DELETE("/classes/:id" , api.DeleteClass)//删除分类
			admin.GET("/users",api.GetAllUser)//获取全部用户
			admin.DELETE("/user/:id",api.DeleteUser)//删除用户
		}
	}
	return r
}