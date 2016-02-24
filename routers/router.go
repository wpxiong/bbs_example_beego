package routers

import (
	"beego-bbs/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.UserController{}, "*:Index")
	beego.Router("/bbs", &controllers.PostController{}, "*:Index")
	beego.AutoRouter(&controllers.PostController{})
	beego.AutoRouter(&controllers.UserController{})
}
