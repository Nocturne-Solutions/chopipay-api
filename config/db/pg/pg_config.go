package pg

import (
	"context"

	"github.com/go-pg/pg/v11"
)

var Db *pg.DB

func InitConnection() error {
	Db = pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User: "chopipay",
		Password: "chopipay",
		Database: "chopipay",
	})

	err := Db.Ping(context.Background())
	return err
}


