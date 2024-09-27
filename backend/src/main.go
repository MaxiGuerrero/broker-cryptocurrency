package main

import (
	"broker-cryptocurrency/backend/src/healthcheck"
	"broker-cryptocurrency/backend/src/system/orchestrator"
	_server "broker-cryptocurrency/backend/src/system/server"
	"runtime"
)

func main() {
	// Run go routines in parallelism based on CPUs of your server
	runtime.GOMAXPROCS(runtime.NumCPU())
	// create server
	server := _server.CreateServer(8080)
	// Initialize features
	initializers := []orchestrator.Initializer{&healthcheck.Healthcheck{}}
	orchestrator.Orchestration(&initializers, *server.Router)
	// start server
	server.StartServer()
}
