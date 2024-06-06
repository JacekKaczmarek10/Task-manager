package main

import (
	"fmt"
	"log"
	"net/http"
	"task-manager/product"

	"github.com/gorilla/mux"
)

func main() {
	LoadAppConfig()

	router := mux.NewRouter().StrictSlash(true)

	RegisterProductRoutes(router)

	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}

func RegisterProductRoutes(router *mux.Router) {
	var muxBase = "/api/products"
	router.HandleFunc(muxBase, product.GetProducts).Methods("GET")
	router.HandleFunc(fmt.Sprintf("%s/{id}", muxBase), product.GetProductById).Methods("GET")
	router.HandleFunc(muxBase, product.CreateProduct).Methods("POST") // Corrected this line
	router.HandleFunc(fmt.Sprintf("%s/{id}", muxBase), product.UpdateProduct).Methods("PUT")
	router.HandleFunc(fmt.Sprintf("%s/{id}", muxBase), product.DeleteProduct).Methods("DELETE")
}
