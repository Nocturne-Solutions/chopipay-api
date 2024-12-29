package pg

import (
	"chopipay/config/server"
	"context"
	"log"

	"github.com/go-pg/pg/v11"
)

var Db *pg.DB

func InitConnection() *pg.DB {
	host := server.GetEnvironment().PgDbHost
	port := server.GetEnvironment().PgDbPort
	user := server.GetEnvironment().PgDbUser
	password := server.GetEnvironment().PgDbPassword
	database := server.GetEnvironment().PgDbName

	Db = pg.Connect(&pg.Options{
		Addr:     host + ":" + port,
		User:     user,
		Password: password,
		Database: database,
	})

	err := Db.Ping(context.Background())
	failOnError(err, "Failed to connect to PostgreSQL")
	return Db
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
