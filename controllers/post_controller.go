package controllers

import (
	"beego-bbs/models/post"
	"github.com/astaxie/beego"
)

type PostController struct {
	beego.Controller
	repository post.PostRepository
}

func (this *PostController) checkUser() bool {
    user := this.GetSession("user")
    if user == nil {
       return false
    }
    return true
}

func (this *PostController) Index() {
    if this.checkUser() == false {
        this.Redirect("/", 302)
        return
    }
    
	posts, _ := this.repository.FindAll()
	this.Data["posts"] = posts
	this.Layout = "layouts/application.html"
	this.TplName = "post/index.html"
}

func (this *PostController) Create() {
    if !(this.checkUser()) {
        this.Redirect("/", 302)
    }
    
	post := post.Post{
		Content: this.GetString("content"),
	}

	err := this.repository.Save(&post)
	flash := beego.NewFlash()
	if err != nil {
		flash.Error("The post could not be saved. Please, try again.")
	} else {
		flash.Notice("The post has been saved.")
	}
	flash.Store(&this.Controller)
	
	this.Redirect("/bbs", 302)
}
