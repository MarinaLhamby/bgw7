package main

import (
	"app/internal/application"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

func main() {
	argsWithProg := os.Args
	if len(argsWithProg) < 1 {
		fmt.Println("You should inform which app you want to run: api or load-job")
	}

	cfg := &application.ConfigApplicationDefault{
		Db: &mysql.Config{
			User:   os.Getenv("DB_USER"),
			Passwd: os.Getenv("DB_PASSWORD"),
			Net:    "tcp",
			Addr:   fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
			DBName: os.Getenv("DB_NAME"),
		},
		Addr: "127.0.0.1:8080",
	}
	var app application.Application
	if argsWithProg[1] == "api" {
		app = application.NewApplicationDefault(cfg)
	} else if argsWithProg[1] == "load-job" {
		app = application.NewLoadApplication(cfg)
	} else {
		fmt.Println("You can only inform api or load-job")
		return
	}

	// - set up
	err := app.SetUp()
	if err != nil {
		fmt.Println(err)
		return
	}
	// - run
	err = app.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
