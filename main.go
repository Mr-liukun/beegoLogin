package main


import (
	_ "beegoLogin/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

)

// 初始化mysql环境
func init() {

	orm.RegisterDriver("mysql", orm.DRMySQL) //注册驱动
	//orm.RegisterModel(new(Model))//注册 model

	// 注册自己使用的数据库
	//orm.RegisterDataBase("qwe", "mysql", "root:123@tcp(localhost:3306)/t7?charset=utf8", 30, 30)

	// 注册默认数据库
	orm.RegisterDataBase("default", "mysql", "root:123@tcp(localhost:3306)/t7?charset=utf8", 30, 30)
}

func main() {
	beego.Run()
}

