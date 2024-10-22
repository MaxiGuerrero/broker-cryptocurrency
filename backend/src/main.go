package main

import (
	"broker-cryptocurrency/backend/src/healthcheck"
	system "broker-cryptocurrency/backend/src/system"
	"runtime"
)

func main() {
	// Run go routines in parallelism based on CPUs of your server
	runtime.GOMAXPROCS(runtime.NumCPU())
	// create server
	server := system.CreateServer(8080)
	// Initialize features
	healthcheck.InitContainer()
	system.PlugInFeatures(getHandlers(), *server.Router)
	// start server
	server.StartServer()
}

func getHandlers() *[]system.Handler {
	// Here, it list all features handlers of your system.
	// Can also remove and add features handlers in the slice.
	return &[]system.Handler{healthcheck.Handler}
}
