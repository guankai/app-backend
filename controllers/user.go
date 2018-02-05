package controllers

import (
	"github.com/astaxie/beego"
	"app-backend/models"
	"encoding/json"
)

type AppUserController struct {
	BaseController
}

// @Description register new user
// @Param body body InputJson true "input json data"
// @router /add [post]
func (c *AppUserController) AddUser() {
	var registerForm models.RegisterForm
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &registerForm)
	if err != nil {
		beego.Debug("parse request error", err)
		c.Data["json"] = models.NewErrorInfo(ErrInputData)
		c.ServeJSON()
		return
	}
	user, err := models.NewUser(&registerForm)
	if err != nil {
		beego.Error("create new user error:", err)
		c.Data["json"] = models.NewErrorInfo(ErrDatabase)
		c.ServeJSON()
		return
	}
	if code, err := user.Insert(); err != nil {
		beego.Error("Insert DB error:", err)
		if code == models.ErrDupRows {
			c.Data["json"] = models.NewErrorInfo(ErrDupUser)
		} else {
			c.Data["json"] = models.NewErrorInfo(ErrDatabase)
		}
		c.ServeJSON()
		return
	}
	c.RetSuccess(user)
	//c.Data["json"] = map[string]interface{}{"code": "0000", "message": "Success", "data": user}
	//c.ServeJSON()
}
