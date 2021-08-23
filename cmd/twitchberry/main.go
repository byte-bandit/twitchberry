package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/klausklapper/twitchberry/internal/api"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			select {
			case s := <-signalChan:
				switch s {
				case syscall.SIGINT, syscall.SIGTERM:
					cancel()
					os.Exit(1)
				case syscall.SIGHUP:
					log.Printf("Got SIGHUP, reloading.")
				}
			case <-ctx.Done():
				log.Printf("Done.")
				os.Exit(1)
			}
		}
	}()

	api.ListenAndServe(ctx)
}
