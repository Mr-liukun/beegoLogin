package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}


// @router /controllers/index [get]
func (c *MainController) Index() {
	c.TplName = "login/login.html"
}

