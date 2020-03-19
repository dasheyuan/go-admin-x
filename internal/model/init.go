package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-admin-x/internal/model/app"
	"go-admin-x/internal/model/sys"
	"go-admin-x/internal/util/conf"
	"go-admin-x/internal/util/orm"
	"log"
	"os"
	"time"
)

func InitDB() {
	var gdb *gorm.DB
	var err error
	var dsn string
	if conf.GetString("DB_DRIVER") == "mysql" {
		dsn = conf.GetString("DB_DEFAULT_DSN")
	} else if conf.GetString("DB_DRIVER") == "sqlite3" {
		dsn = conf.GetString("DB_PATH")
	}
	gdb, err = gorm.Open(conf.GetString("DB_DRIVER"), dsn)
	if err != nil {
		panic(err)
	}
	gdb.SingularTable(true)
	if conf.GetBool("DB_ORM_DEBUG") {
		gdb.LogMode(true)
		gdb.SetLogger(log.New(os.Stdout, "\r\n", 0))
	}
	gdb.DB().SetMaxIdleConns(conf.GetInt("ORM_MAX_IDLE_CONNS"))
	gdb.DB().SetMaxOpenConns(conf.GetInt("ORM_MAX_OPEN_CONNS"))
	gdb.DB().SetConnMaxLifetime(time.Duration(conf.GetInt("ORM_MAX_LIFETIME")) * time.Second)
	orm.DB = gdb
}

func Migration() {
	fmt.Println(orm.DB.AutoMigrate(new(sys.Menu)).Error)
	fmt.Println(orm.DB.AutoMigrate(new(sys.Admin)).Error)
	fmt.Println(orm.DB.AutoMigrate(new(sys.RoleMenu)).Error)
	fmt.Println(orm.DB.AutoMigrate(new(sys.Role)).Error)
	fmt.Println(orm.DB.AutoMigrate(new(sys.AdminRole)).Error)
	fmt.Println(orm.DB.AutoMigrate(new(app.Version)).Error)
}
