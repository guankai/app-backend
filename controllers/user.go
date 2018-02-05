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

// @Description user login
// @param login body InputJson true "login json"
// @Success 200 {string}
// @router /login [post]
func (c *AppUserController) Login() {
	beego.Info("app login start...")
	loginForm := models.LoginForm{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &loginForm); err != nil {
		c.RetError(errInputData)
		return
	}
	user := models.AppUser{}
	if code, err := user.FindByPhone(loginForm.Phone); err != nil {
		beego.Error("FindByPhone", err)
		if code == models.ErrNotFound {
			c.RetError(errNoUser)
		} else {
			c.RetError(errDatabase)
		}
		return
	}
	if ok, err := user.CheckPass(loginForm.Password); err != nil {
		beego.Error("CheckUserPassword", err)
		c.RetError(errDatabase)
		return
	}else if !ok{
		c.RetError(errPass)
		return
	}
	user.ClearPass()
	c.RetSuccess(&user)
}
