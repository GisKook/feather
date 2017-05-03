package main

import (
	"fmt"
	"github.com/giskook/feather/conf"
	"github.com/giskook/feather/db_socket"
	"github.com/giskook/feather/feather_worker"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// read configuration
	configuration, _ := conf.ReadConfig("./conf.json")

	db_socket.NewDbSocket(configuration.DB)
	worker := feather_worker.NewFeatherWorker()
	go worker.Do()

	// catchs system signal
	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Signal: ", <-chSig)
}
