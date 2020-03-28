package service

import (
	"MyBlog/cache"
	"MyBlog/model"
	"MyBlog/serializer"
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"time"
)

type UserService struct {
	Username        string    `form:"username" json:"username" binding:"required,min=5,max=30"`
	Password  		string    `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string 	  `form:"password_confirm" json:"password_confirm"`
}
//设置session
func (service *UserService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.Id)
	_ = s.Save()
}
//登录
func (service *UserService) SignIn(c *gin.Context) serializer.Response  {
	var user model.User

	if err := model.DB.Where("username = ? AND password = ?", service.Username, service.Password).First(&user).Error; err != nil {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	// 设置session
	service.setSession(c, user)
	//存放信息到Redis中
	jsonStr,_ :=json.Marshal(serializer.BuildUser(user))
	cache.RedisClient.Set(user.Id,jsonStr,0)

	return serializer.BuildUserResponse(user)
}
//注册
func (service *UserService) SignUp()  serializer.Response{
	//local, _ := time.LoadLocation("Local")
	user := model.User{
		Id: model.GetNewUserId("02"),
		Username:       service.Username,
		PasswordDigest: service.Password,
		Created: time.Now(),
		Updated: time.Now(),
	}

	// 表单验证
	//if err := service.Valid(); err != nil {
	//	return *err
	//}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.ParamErr("注册失败", err)
	}

	return serializer.BuildUserResponse(user)
}

//验证
func (service *UserService) Valid() *serializer.Response{
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Code: 40001,
			Msg:  "两次输入的密码不相同",
		}
	}

	return nil
}
//获取所有用户
func (service *UserService) GetAllUser() serializer.Response  {
	if list, err := model.GetAllUser(); err != nil {
		return serializer.DBErr("连接数据库失败",err)
	} else {
		return serializer.Response{
			Code:  200,
			Data:  list,
		}
	}
}

//删除用户
func (service *UserService) DeleteUser(id string) serializer.Response {
	//var count int
	if id == "010000000001"{
		return serializer.ParamErr("管理员用户无法删除",nil)
	}else {
		//将评论的用户换为注销用户
		model.DB.Model(model.Comments{}).Where("uid=?",id).Updates(map[string]interface{}{"uid":"000000000000"})
		model.DB.Model(model.Comments{}).Where("superiorUser=?",id).Updates(map[string]interface{}{"superiorUser":"000000000000"})
		/*if re :=model.DB.Model(model.Comments{}).Where("uid=?",id).Update("uid=?","000000000000");re.Error != nil{
			return serializer.DBErr("连接数据失败",re.Error)
		}
		if re :=model.DB.Model(model.Comments{}).Where("superiorUser=?",id).Update("superiorUser=?","000000000000");re.Error != nil{
			return serializer.DBErr("连接数据失败",re.Error)
		}*/

		//删除操作
		if eRow , err :=model.DeleteUser(model.User{Id:       id,});err!=nil{
			return serializer.DBErr("连接数据失败",err)
		}else {
			return serializer.Response{
				Code:  200,
				Data:  eRow,
			}
		}
	}

}