package main

import (
	"fmt"
	"time"
	"strings"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "beego-bbs/routers"
	"github.com/astaxie/beego/logs"
)

func init() {
	user := beego.AppConfig.String("mysqluser")
	pass := beego.AppConfig.String("mysqlpass")
	pass = ""
	host := beego.AppConfig.String("mysqlurls")
	db := beego.AppConfig.String("mysqldb")
	str_v := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8", user, pass, host, db)
	fmt.Println(str_v)
    datasource := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8", user, pass, host, db)
	orm.RegisterDataBase("default", "mysql", datasource, 30)

    log := logs.NewLogger(1000)
	log.SetLevel(logs.LevelDebug)
	log.SetLogger("console", fmt.Sprintf("{\"level\":%d}", logs.LevelNotice))
	log.Debug("%s", "this is a debug message")
	log.Informational("%s", "this is an informational message")
	log.Notice("%s", "this is a notice message")
	log.Flush()

	
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		panic(err)
	}
    fmt.Scanln()
	log.Close()

	beego.AddFuncMap("dateformatJst", func(in time.Time) string {
		in = in.Add(time.Duration(9) * time.Hour)
		return in.Format("2006-01-02 15:04:05")
	})
	
	beego.AddFuncMap("nl2br", func(in string) string {
		return strings.Replace(in, "\n", "<br>", -1)
	})
}



func main() {
	beego.Run()
}
