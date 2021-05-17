package main

import (
	"books-list/controllers"
	"books-list/driver"
	"books-list/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db := driver.ConnectDB()
	controller := controllers.Controller{}

	router := mux.NewRouter()

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")

	fmt.Println("Server is running at port 8000")
	err := http.ListenAndServe(":8000", router)
	utils.LogFatal(err)
}
