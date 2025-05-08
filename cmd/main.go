package main

import (
	"golang.org/x/net/context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"restfullApi/pkg/api"
	"syscall"
)

var (
	osInterrupt       chan os.Signal
	listenerInterrupt chan error
)

func init() {
	osInterrupt = make(chan os.Signal)
	listenerInterrupt = make(chan error)
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	log.Printf("start application service, listener address: %s:%s", "", "")

	srv := http.Server{
		Addr:    ":9090",
		Handler: api.Run(),
	}
	go func() {
		defer log.Println("listener has been stopped")

		log.Println("starting listener...")
		if err := srv.ListenAndServe(); err != nil {
			listenerInterrupt <- err
		}
		log.Println("listener has been started")
	}()

	signal.Notify(osInterrupt, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	select {
	case <-osInterrupt:
		log.Println("OS interrupt received: shutting down server gracefully....")
		_ = srv.Shutdown(ctx)
	case err := <-listenerInterrupt:
		log.Printf("Server listener encountered an error:%v shutting down....", err)
	}

}
