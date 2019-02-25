package login

import "github.com/astaxie/beego/orm"

type User struct {
	Username string
	Password string
}

func LoginModel(username,password string) *User {
	o := orm.NewOrm()

	//o.Using("qwe") // 选择数据库，不设置使用默认数据库
	var user *User

	o.Raw("select username, password from " +
		"user where username = ? and password = ?",username, password).QueryRow(&user)
	return user
}
