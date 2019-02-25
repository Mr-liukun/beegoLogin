package routers

import (
	"beegoLogin/controllers"
	"beegoLogin/controllers/login"

	"github.com/astaxie/beego"
)

func init() {



	beego.Include(&login.LoginController{})
	beego.Include(&controllers.MainController{})

	//beego.Router("/index",&controllers.MainController{},"get:Index")

   // beego.Router("/login", &login.LoginController{},"post:Login")
}
