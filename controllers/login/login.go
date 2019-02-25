package login

import (
	"beegoLogin/models/login"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

// @router /login/login [post]
func (lc *LoginController) Login() {
	username := lc.GetString("username")
	password := lc.GetString("password")

	user := login.LoginModel(username,password)
	//fmt.Println(*user)
	if(user == nil){
		lc.TplName = "login/fail.html"
	} else {
	//	// 返回多种不同数据的json
	//	//stu := Student{Id:10,Name:"qwe"}
	//	//stu1 := Student{Id:8,Name:"awsd"}
	//	//
	//	//st := [...]Student{stu,stu1}
	//	//
	//	//ss := [...]interface{}{st,user}
	//	//
	//	//lc.Data["json"] = ss
	//	//lc.SetSession("user",user)
		lc.Data["username"] = user.Username
		lc.Data["password"] = user.Password
		lc.TplName = "login/success.html"
		//lc.ServeJSON()
	}
}


