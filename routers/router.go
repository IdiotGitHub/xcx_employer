package routers

import (
	"github.com/astaxie/beego"
	"xcx_employer/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user", &controllers.UserInfoController{}, "get:GetOne")
    beego.Router("/test", &controllers.UserInfoController{}, "get:Test")
}
