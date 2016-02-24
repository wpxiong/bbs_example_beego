package user

import (
	"time"
)

type User struct {
	Userid int64 `orm:"pk;auto"`
	Username string `orm:"type(text)"`
	Password string `orm:"type(text)"`
	Email string `orm:"type(text)"`
	Role  string `orm:"type(text)"`
	LastLogin time.Time
	Created time.Time
	Modified time.Time
}
