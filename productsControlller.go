package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (app *Application) GetProducts(w http.ResponseWriter, r *http.Request) {
	products := app.db.Find(&Products{})
	parsedProduct, _ := json.Marshal(products)
	w.Write(parsedProduct)
}

func (app *Application) GetProductsById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Id, err := strconv.Atoi(params["Id"])
	if err != nil {
		app.serverError(w, err)
		return
	}
	products := app.db.First(&Products{}, Id)
	parsedProduct, _ := json.Marshal(products)
	w.Write(parsedProduct)
}

func (app *Application) PostProduct(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var product Products
	err := decoder.Decode(&product)
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.db.Create(&product)
	w.Write([]byte(fmt.Sprintf("Created %d", product.ID)))
}

func (app *Application) PutProduct(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var product Products
	err := decoder.Decode(&product)
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.db.Update(&product)
	w.Write([]byte(fmt.Sprintf("Updated %d", product.ID)))
}

func (app *Application) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	Id, err := strconv.Atoi(params["Id"])
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.db.Delete(&Products{}, Id)
	w.Write([]byte(fmt.Sprintf("Deleted %d", Id)))
}
