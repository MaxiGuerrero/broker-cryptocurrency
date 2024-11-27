package main

import (
	"backend/src/authentication"
	"backend/src/healthcheck"
	system "backend/src/system"
	"runtime"
)

func main() {
	// Run go routines in parallelism based on CPUs of your server
	runtime.GOMAXPROCS(runtime.NumCPU())
	// create server
	server := system.CreateServer(8080)
	// Initalize dependencies
	initAllDependencies()
	// Registering all handlers of each module in the system
	system.RegisterHandlers(getHandlers(), *server.Router)
	// start server
	server.StartServer()
}

func getHandlers() *[]system.Handler {
	// Here, it list all features handlers of the system.
	// Can also remove and add features handlers in the slice.
	return &[]system.Handler{
		healthcheck.Handler,
		authentication.Handler,
	}
}

func initAllDependencies() {
	// Here, it call all methods init dependencies of the system.
	healthcheck.InitContainerDependency()
	authentication.InitContainerDependency()
}
