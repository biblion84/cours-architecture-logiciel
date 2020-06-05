package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jinzhu/gorm"
)

type Application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	db       *gorm.DB
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	err := ConnectDatabase()
	if err != nil {
		panic("Could not connect to database")
	}
	app := &Application{
		errorLog: errorLog,
		infoLog:  infoLog,
		db:       Db,
	}
	srv := &http.Server{
		Handler:      app.routes(),
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	infoLog.Printf("Starting server on :8080")
	log.Fatal(srv.ListenAndServe())

}
