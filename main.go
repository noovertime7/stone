package main

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/noovertime7/stone/router"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	lib.InitModule("./conf/dev/", []string{"base", "mysql"})
	defer lib.Destroy()
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
