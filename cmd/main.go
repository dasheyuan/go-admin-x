package main

import (
	"flag"
	"go-admin-x/internal/app"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	flag.Parse()
	//log.Init(nil) // debug flag: log.dir={path}
	//defer log.Close()
	//log.Info("ai-block start")
	//paladin.Init(apollo.PaladinDriverApollo)
	//_, closeFunc, err := di.InitApp()
	//if err != nil {
	//	panic(err)
	//}
	app.Run()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		//log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			//closeFunc()
			//log.Info("ai-block exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
