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
// @Success 200 {string}
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
	userOut := models.UserOut{Id: user.ID, Phone: user.Phone, Name: user.Name}
	c.RetSuccess(&userOut)
}

// @Description delete user
// @Param userId path int true "user id"
// @Success 200 {string}
// @router /:userId/delete [delete]
func (c *AppUserController) DeleteUser() {
	appId, errParse := c.GetInt(":userId")
	if errParse != nil {
		c.RetError(errInputData)
		return
	}
	user := new(models.AppUser)
	_, err := user.FindById(appId)
	if err != nil {
		c.RetError(errDatabase)
		return
	}
	_, errDel := user.DeleteUser()
	if errDel != nil {
		c.RetError(errDatabase)
		return
	}
	c.RetSuccess("delete user success")
}
