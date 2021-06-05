package main

import (
	"books-list/driver"
	"books-list/handlers"
	"books-list/utils"
	"fmt"
	"net/http"

	gh "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

func init() {
	err := gotenv.Load("../.env")
	utils.LogFatal(err)
}

func main() {
	db := driver.ConnectDB()
	h := handlers.Handler{}

	router := mux.NewRouter()

	router.HandleFunc("/books", h.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", h.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", h.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", h.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", h.RemoveBook(db)).Methods("DELETE")

	fmt.Println("Server is running at port 8000")
	err := http.ListenAndServe(":8000", gh.CORS(gh.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		gh.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		gh.AllowedOrigins([]string{"*"}))(router))
	utils.LogFatal(err)
}
