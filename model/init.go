package model

import (
	"MyBlog/util"
	"github.com/jinzhu/gorm"
	"time"
)

//DB数据库连接单例
var DB *gorm.DB

func Database(connString string) {
	db, err := gorm.Open("mysql",connString)
	db.LogMode(true)
	//Error
	if err != nil {
		util.Log().Panic("连接数据库不成功",err)
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(10)
	//打开
	db.DB().SetMaxOpenConns(50)
	//超时
	db.DB().SetConnMaxLifetime(time.Second*30)

	DB = db

	migration()
}