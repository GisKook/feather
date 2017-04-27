package main

import (
	"fmt"
	"github.com/giskook/feather/conf"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// read configuration
	configuration, err := conf.ReadConfig("./conf.json")

	// catchs system signal
	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Signal: ", <-chSig)
	redis_socket.Close()
	mq_socket.Stop()
	db.GetDBSocket().Listener.UnlistenAll()
}
