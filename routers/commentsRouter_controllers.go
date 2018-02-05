package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["app-backend/controllers:AppUserController"] = append(beego.GlobalControllerRouter["app-backend/controllers:AppUserController"],
		beego.ControllerComments{
			Method: "AddUser",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["app-backend/controllers:AppUserController"] = append(beego.GlobalControllerRouter["app-backend/controllers:AppUserController"],
		beego.ControllerComments{
			Method: "DeleteUser",
			Router: `/:userId/delete`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["app-backend/controllers:ObjectController"] = append(beego.GlobalControllerRouter["app-backend/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["app-backend/controllers:ObjectController"] = append(beego.GlobalControllerRouter["app-backend/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["app-backend/controllers:ObjectController"] = append(beego.GlobalControllerRouter["app-backend/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["app-backend/controllers:ObjectController"] = append(beego.GlobalControllerRouter["app-backend/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["app-backend/controllers:ObjectController"] = append(beego.GlobalControllerRouter["app-backend/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
