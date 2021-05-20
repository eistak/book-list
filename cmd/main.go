package main

import (
	"books-list/driver"
	"books-list/handlers"
	"books-list/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

func init() {
	err := gotenv.Load("../.env")
	utils.LogFatal(err)
}

func main() {
	db := driver.ConnectDB()
	handlers := handlers.Handler{}

	router := mux.NewRouter()

	router.HandleFunc("/books", handlers.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", handlers.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", handlers.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", handlers.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", handlers.RemoveBook(db)).Methods("DELETE")

	fmt.Println("Server is running at port 8000")
	err := http.ListenAndServe(":8000", router)
	utils.LogFatal(err)
}
