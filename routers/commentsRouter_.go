package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["beegoLogin/controllers:MainController"] = append(beego.GlobalControllerRouter["beegoLogin/controllers:MainController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/controllers/index`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
