package main

import (
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func (app *Application) routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/Products", app.GetProducts)
	r.HandleFunc("/Products/{Id:[0-9]+}", app.GetProductsById)
	r.HandleFunc("/Products", app.PostProduct).Methods(http.MethodPost)
	r.HandleFunc("/Products", app.PutProduct).Methods(http.MethodPut)
	r.HandleFunc("/Products/{Id:[0-9]+}", app.DeleteProduct).Methods(http.MethodDelete)
	return r
}
