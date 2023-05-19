package main

import (
	_ "github.com/abelz123456/celestial-api/docs"
	"github.com/abelz123456/celestial-api/package/log"
	"github.com/abelz123456/celestial-api/package/manager"
	"github.com/abelz123456/celestial-api/routes"
)

// @title Celestialsoftware API
// @version v1.0
// @BasePath /
func main() {
	manager, err := manager.Init(".")
	if err != nil {
		log.NewLog().Panic(err, "", nil)
	}

	routes.LoadRoute(*manager)

	log.NewLog().Info("Http Server loaded", map[string]string{"address": manager.Server.HttpServer.Addr})
	manager.Server.HttpServer.ListenAndServe()
}
