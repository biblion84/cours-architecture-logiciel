package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	db       *gorm.DB
}

var Db *gorm.DB

const connectionString = "Your connection string"

func ConnectDatabase() error {
	var err error
	Db, err = gorm.Open("postgres", connectionString)
	if err != nil {
		return fmt.Errorf("error in connectDatabase(): %v", err)
	}
	Db.AutoMigrate(&Products{})
	return err
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
