package config

import (
	"github.com/go-pg/pg/v9"
	"log"
	"os"

	controllers "github.com/flyq/minigram/controllers"
)

// Connecting to db
func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "ubuntu",
		Password: "123456",
		Addr:     "localhost:5432",
		Database: "middleware_hn",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	controllers.CreateUserTable(db)
	controllers.CreateEleRecTable(db)
	controllers.CreateBloRecTable(db)
	controllers.InitiateDB(db)
	return db
}
