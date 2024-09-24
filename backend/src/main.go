package main

import (
	"broker-cryptocurrency/backend/src/healthcheck"
	_server "broker-cryptocurrency/backend/src/server"
	"runtime"
)

func main() {
	// Run go routines in parallelism based on CPUs of your server
	runtime.GOMAXPROCS(runtime.NumCPU())
	// create server
	server := _server.CreateServer(8080)
	// register routes
	healthcheck.RegisterRoutes(*server.Router)
	// start server
	server.StartServer()
}
