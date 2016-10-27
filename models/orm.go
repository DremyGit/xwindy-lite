package models

import (
	"errors"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
)

var (
	o orm.Ormer
)

// DatabaseCheck check the status of database connection
type DatabaseCheck struct{}

func (dc *DatabaseCheck) isConnect() bool {
	return true
}

// Check check
func (dc *DatabaseCheck) Check() error {
	if !dc.isConnect() {
		return errors.New("Can't connect database")
	}
	return nil
}

func init() {
	iniconf, err := config.NewConfig("ini", "./conf/app.conf")
	if err != nil {
		panic("Config file not found")
	}

	user := iniconf.String("database::user")
	pass := iniconf.String("database::pass")
	db := iniconf.String("database::database")
	orm.RegisterDataBase("default", "mysql",
		user+":"+pass+"@/"+db+"?charset=utf8&loc=Asia%2FShanghai")
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)

	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
	}

	o = orm.NewOrm()
	toolbox.AddHealthCheck("database", &DatabaseCheck{})
}
