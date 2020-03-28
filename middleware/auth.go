package middleware

import (
	"MyBlog/model"
	"MyBlog/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("user_id")
		if uid != nil {
			user, err := model.GetUser(uid)
			if err == nil {
				c.Set("user", &user)
			}
		}
		c.Next()
	}
}

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*model.User); ok {
				c.Next()
				return
			}
		}

		c.JSON(200, serializer.CheckLogin())
		c.Abort()
	}
}

// AdminRequired 需要管理员权限
func AdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if admin, ok := user.(*model.User); ok && admin.Id[0:2] == "01"  {
				c.Next()
				return
			}

			c.JSON(200, serializer.NoRightErr())
			c.Abort()
		}

		c.JSON(200, serializer.CheckLogin())
		c.Abort()
	}
}