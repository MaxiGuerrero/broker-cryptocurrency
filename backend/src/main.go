package main

import (
	"backend/src/authentication"
	"backend/src/healthcheck"
	system "backend/src/system"
	"backend/src/system/database"
	"backend/src/wallet"
	"context"
	"runtime"
)

func main() {
	ctx := context.Background()
	// Init Dotenv
	system.InitLoadEnv()
	// Run go routines in parallelism based on CPUs of your server
	runtime.GOMAXPROCS(runtime.NumCPU())
	// create server
	server := system.CreateServer(8080)
	// Init database
	db := database.NewDatabase(ctx)
	// Initalize dependencies
	initAllDependencies(db)
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
		wallet.Handler,
	}
}

func initAllDependencies(db *database.Database) {
	// Here, it call all methods init dependencies of the system.
	healthcheck.InitContainerDependency()
	authentication.InitContainerDependency(db)
	wallet.InitContainerDependency(db)
}
