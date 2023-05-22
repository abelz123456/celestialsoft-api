package main

import (
	"github.com/abelz123456/celestial-api/entity"
	"github.com/abelz123456/celestial-api/package/config"
	"github.com/abelz123456/celestial-api/package/database"
	"github.com/abelz123456/celestial-api/package/log"
)

var registeredModels []interface{} = []interface{}{
	entity.Bank{},
	entity.PermissionPolicyUser{},
	entity.LocalFile{},
	entity.Rajaongkir{},
	entity.EmailSent{},
}

func main() {
	var (
		conf config.Config = config.Init(".")
		log  log.Log       = log.NewLog()
	)

	db, err := database.NewDatabase(conf)
	log.PanicOnError(err, "Database migration aborted", nil)

	if db.Driver == database.Mongo {
		log.Info("Database migration aborted, because the database used does not require this", map[string]string{"dbDriver": db.Driver.String()})
		return
	}

	for _, model := range registeredModels {
		log.Info("Database migration begins", nil)
		log.PanicOnError(db.Sql.AutoMigrate(&model), "", nil)
	}
	log.Info("Database migration complete", nil)
}
