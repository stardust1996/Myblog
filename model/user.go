package model

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	//gorm.Model
	Id             string    `gorm:"column:id"`
	Username       string    `gorm:"column:username"`
	PasswordDigest string    `gorm:"column:password"`
	Created        time.Time `gorm:"column:created"`
	Updated        time.Time `gorm:"column:updated"`
}

const (
	//加密难度
	PassWordCost = 12
)
//查找全部用户
func GetAllUser() (users []User,err error)  {
	err = DB.Where("id <> ?", "000000000000").Order("id").Find(&users).Error
	return
}

//按照ID获取用户信息
func GetUser(ID interface{}) (User , error){
	var user User
	result := DB.First(&user,ID)
	return user,result.Error
}

//设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

//删除用户
func DeleteUser(user User) (int64, error) {
	result := DB.Delete(&user)
	return result.RowsAffected,result.Error
}