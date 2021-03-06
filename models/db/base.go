package db

import (
	"sync"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
	"fmt"
	"github.com/astaxie/beego/logs"
	//"database/sql"
)

var globalOrm orm.Ormer
// Database is an interface of different databases
type Database interface {
	// Name() returns the name of database
	Name() string
	// String() returns the details of the database
	String() string
	// Register registers the databse which will be used
	Register(alias ...string) error
}

func InitDatabase() {
	database, err := getDatabase()
	if err != nil {
		panic(err)
	}

	logs.Info("initializing database: %s", database.String())
	if err := database.Register(); err != nil {
		panic(err)
	}
	globalOrm = orm.NewOrm()
}

func getDatabase() (db Database, err error) {
	switch strings.ToLower(beego.AppConfig.String("database")) {
	case "", "mysql":
		host, port, user, pwd, database := getMySQLConnInfo()
		db = NewMySQL(host, port, user, pwd, database)
	default:
		err = fmt.Errorf("invalid database: %s", beego.AppConfig.String("database"))
	}

	return
}

func getMySQLConnInfo() (host, port, username, password, database string) {
	host = beego.AppConfig.String("mysql::url")
	port = beego.AppConfig.String("mysql::port")
	username = beego.AppConfig.String("mysql::user")
	password = beego.AppConfig.String("mysql::pwd")
	database = beego.AppConfig.String("mysql::db")
	if len(database) == 0 {
		database = "app-backend"
	}
	return
}

var once sync.Once

func GetOrmer() orm.Ormer {
	once.Do(func() {
		globalOrm = orm.NewOrm()
	})
	return globalOrm
}