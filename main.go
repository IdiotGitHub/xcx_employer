package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/go-sql-driver/mysql"
	_ "xcx_employer/routers"
)

func main() {
	err := orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1)/xiaochengxu")
	//处理跨域问题
	//InsertFilter是提供一个过滤函数
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		//允许访问所有源
		AllowAllOrigins: true,
		//可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
		//其中Options跨域复杂请求预检
		AllowMethods:   []string{"*"},
		//指的是允许的Header的种类
		AllowHeaders: 	[]string{"*"},
		//公开的HTTP标头列表
		ExposeHeaders:	[]string{"Content-Length"},
		//如果设置，则允许共享身份验证凭据，例如cookie
		AllowCredentials: true,
	}))
	if err != nil {
		fmt.Println("connect database error")
		return
	}
	beego.Run()
}

