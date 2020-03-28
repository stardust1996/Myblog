package main

import (
	"MyBlog/conf"
	"MyBlog/server"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//读取配置
	conf.Init()

	r := server.NewRouter()
	_ = r.Run(":8080")
}