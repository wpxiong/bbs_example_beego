package controllers

import (
	"beego-bbs/models/user"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
	repository user.UserRepository
}

func (this *UserController) Index() {
	this.Layout = "layouts/application.html"
	this.TplName = "login/index.html"
}

func (this *UserController) Logout() {
    this.DelSession("user")
	this.Layout = "layouts/application.html"
	this.TplName = "login/index.html"
}

func (this *UserController) Login() {
	username := this.Input().Get("username")
	password := this.Input().Get("password")
	this.Data = make(map[interface{}]interface{})
    if len(username) <= 4  {
        this.Data["usernameErr"] = "username must be 4 letters at least."
    }
    if len(password) <= 6 {
        this.Data["passwordErr"] = "password must be 6 letters at least."
    }
    size := len(this.Data)
    beego.Debug(size)
    if size > 0 {
          this.Index()
          return
    }else{
        users, _ := this.repository.FindByUserName(username)
        beego.Debug(users)
        if len(users) > 0 {
            this.SetSession("user",users[0])
            this.Redirect("/bbs",302)
        }else{
            this.Data["passErr"] = "username or password is not correct."
            this.Index()
            return
        }
    }
}
