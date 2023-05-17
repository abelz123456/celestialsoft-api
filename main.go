package main

import (
	"github.com/abelz123456/celestial-api/package/log"
	"github.com/abelz123456/celestial-api/package/manager"
	"github.com/abelz123456/celestial-api/routes"
)

func main() {
	manager, err := manager.Init(".")
	if err != nil {
		log.NewLog().Panic(err, "", nil)
	}

	routes.LoadRoute(*manager)

	manager.Server.HttpServer.ListenAndServe()
}
