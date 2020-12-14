package infra

import (
	"github.com/ablarry/golang-bootcamp-2020/pkg/config"

	"github.com/jmoiron/sqlx"
	// Driver BD
	_ "github.com/lib/pq"
)

// InitDB initialize connection to DB
func InitDB(c *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect(c.DB.Driver, c.ConnectURL())
	if err != nil {
		return nil, err
	}
	// ping DB
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(c.DB.MaxIdleConns)
	db.SetMaxOpenConns(c.DB.MaxOpenConns)
	return db, nil

}
