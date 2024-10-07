package pg

import (
	"context"
	"log"

	"github.com/go-pg/pg/v11"
)

var Db *pg.DB

func InitConnection(envVars map[string]string) {
	host := envVars["PG_DB_HOST"]
	port := envVars["PG_DB_PORT"]
	user := envVars["PG_DB_USER"]
	password := envVars["PG_DB_PASSWORD"]
	database := envVars["PG_DB_NAME"]

	if host == "" || port == "" || user == "" || password == "" || database == "" {
		log.Panic("Missing environment variables for postgres database connection")
	}

	Db = pg.Connect(&pg.Options{
		Addr:     host + ":" + port,
		User: 	  user,
		Password: password,
		Database: database,
	})

	err := Db.Ping(context.Background())
	failOnError(err, "Failed to connect to PostgreSQL")
}

func CloseConnection() {
	err := Db.Close(context.Background())
	failOnError(err, "Failed to close connection to PostgreSQL")
}

func failOnError(err error, msg string) {
	if err != nil {
	  log.Panicf("%s: %s", msg, err)
	}
}


