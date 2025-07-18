package application

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

// Application is an interface that represents an application.
type Application interface {
	// Run runs the application.
	Run() (err error)
	// SetUp sets up the application.
	SetUp() (err error)
}

// ConfigApplicationDefault is the configuration for NewApplicationDefault.
type ConfigApplicationDefault struct {
	// Db is the database configuration.
	Db *mysql.Config
	// Addr is the server address.
	Addr string
}

// ApplicationDefault is an implementation of the Application interface.
type ApplicationConfig struct {
	// cfgDb is the database configuration.
	cfgDb *mysql.Config
	// cfgAddr is the server address.
	db *sql.DB
}

func (a *ApplicationConfig) SetUpDb() (err error) {
	// dependencies
	// - db: init
	db, err := sql.Open("mysql", a.cfgDb.FormatDSN())
	fmt.Println(a.cfgDb)
	if err != nil {
		fmt.Println("error opening connection")
		return
	}
	// - db: ping
	err = db.Ping()
	if err != nil {
		fmt.Println("error connecting to DB")
		return
	}
	a.db = db
	return nil
}
