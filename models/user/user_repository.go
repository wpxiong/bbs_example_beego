package user

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type UserRepository struct {

}

func init() {
	orm.RegisterModel(new(User))
}

func (this *UserRepository) InitData()  {
    user1 := new(User)
	user1.Username = "wp.xiong"
	user1.Password = "password"
	user1.Email = "wp.xiong@gmail.com"
	user1.Role = "Admin"
	user1.LastLogin = time.Now()
	user1.Created = user1.LastLogin
	user1.Modified = user1.LastLogin
	this.Save(user1)
}

func (this *UserRepository) FindAll() ([]*User, error) {
	o := orm.NewOrm()
	var Users []*User
	_, err := o.QueryTable("User").OrderBy("-Userid").All(&Users)
	return Users, err
}

func (this *UserRepository) FindById(id int64) (*User, error) {
	o := orm.NewOrm()
	User := User{Userid: id}
	err := o.Read(&User)
	return &User, err
}


func (this *UserRepository) FindByUserName(username string)  ([]*User, error) {
	o := orm.NewOrm()
	var Users []*User
	_, err := o.QueryTable("User").Filter("Username", username).All(&Users)
	return Users, err
}


func (this *UserRepository) Save(p *User) error {
	var err error
	o := orm.NewOrm()
	now := time.Now()
	
	if p.Userid != 0 {
		err = o.Read(p)
		if err != nil {
			p.Modified = now
			_, err = o.Update(p)
		}
	} else {
		p.Created = now
		p.Modified = now
		_, err = o.Insert(p)
	}
	
	return err
}

func (this *UserRepository) Delete(p *User) error {
	o := orm.NewOrm()
	_, err := o.Delete(p)
	return err
}
