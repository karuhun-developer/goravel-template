package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/goravel/framework/facades"

	"karuhundeveloper.com/gostarterkit/bootstrap"
)

func main() {
	// This bootstraps the framework and gets it ready for use.
	bootstrap.Boot()

	// Create a channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start http server by facades.Route().
	go func() {
		if err := facades.Route().Run(); err != nil {
			facades.Log().Errorf("Route Run error: %v", err)
		}
	}()

	// Start queue server by facades.Queue().
	go func() {
		if err := facades.Queue().Worker().Run(); err != nil {
			facades.Log().Errorf("Queue run error: %v", err)
		}
	}()

	// Start Scheduler server by facades.Schedule().
	go facades.Schedule().Run()

	// Listen for the OS signal
	go func() {
		<-quit
		if err := facades.Route().Shutdown(); err != nil {
			facades.Log().Errorf("Route Shutdown error: %v", err)
		}
		if err := facades.Queue().Worker().Shutdown(); err != nil {
			facades.Log().Errorf("Queue Shutdown error: %v", err)
		}
		if err := facades.Schedule().Shutdown(); err != nil {
			facades.Log().Errorf("Schedule Shutdown error: %v", err)
		}

		os.Exit(0)
	}()

	select {}
}
