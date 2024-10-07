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
	initializers := []system.Route{&healthcheck.Healthcheck{}}
	system.RegisterRoutes(&initializers, *server.Router)
	// start server
	server.StartServer()
}
