package utils

import (
	"database/sql/driver"
	"fmt"
	"os"

	"github.com/DATA-DOG/go-txdb"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func Init() driver.Connector {
	cfg := mysql.Config{
		User:   os.Getenv("DB_ROOT_USER"),
		Passwd: os.Getenv("DB_ROOT_PASSWORD"),
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		DBName: os.Getenv("DB_NAME"),
	}
	// cfg := mysql.Config{
	// 	User:   "root",
	// 	Passwd: "rootpassword",
	// 	Net:    "tcp",
	// 	Addr:   "localhost:3307",
	// 	DBName: "fantasy_products_test",
	// }
	return txdb.New("mysql", cfg.FormatDSN())
}
