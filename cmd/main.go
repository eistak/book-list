package main

import (
	"books-list/api"
	"books-list/tools"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

func init() {
	err := gotenv.Load("../.env")
	tools.LogFatal(err)
}

func main() {
	db := tools.ConnectDB()
	h := api.Api{}

	router := mux.NewRouter()

	router.HandleFunc("/books", h.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", h.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", h.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", h.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", h.RemoveBook(db)).Methods("DELETE")

	fmt.Println("Server is running at port 8000")
	err := http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(router))
	tools.LogFatal(err)
}
