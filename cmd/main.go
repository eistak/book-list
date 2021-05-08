package main

import (
	"books-list/controllers"
	"books-list/driver"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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
	log.Fatal(http.ListenAndServe(":8000", router))
}

//func getBooks(w http.ResponseWriter, r *http.Request) {
//	books = []Book{}
//	db.Find(&books)
//
//	json.NewEncoder(w).Encode(books)
//}

//func getBook(w http.ResponseWriter, r *http.Request) {
//	var book Book
//	params := mux.Vars(r)
//	book.ID = params["id"]
//	db.First(&book)
//
//	json.NewEncoder(w).Encode(book)
//}

//func addBook(w http.ResponseWriter, r *http.Request) {
//	var book Book
//
//	json.NewDecoder(r.Body).Decode(&book)
//	result := db.Create(&book)
//	if result.Error != nil {
//		log.Println(result.Error)
//	}
//	log.Println(result.RowsAffected)
//	fmt.Println("added.")
//
//	json.NewEncoder(w).Encode(book.ID)
//}
//
//func updateBook(w http.ResponseWriter, r *http.Request) {
//	var book Book
//
//	json.NewDecoder(r.Body).Decode(&book)
//	db.Save(&book)
//
//	fmt.Println(book.ID)
//	log.Println("updated.")
//	json.NewEncoder(w).Encode(book.ID)
//}
//
//func removeBook(w http.ResponseWriter, r *http.Request) {
//	var book Book
//	params := mux.Vars(r)
//	book.ID = params["id"]
//	log.Println(book.ID)
//	result := db.Delete(&book, book.ID)
//	if result.Error != nil {
//		log.Println(result.Error)
//	}
//	log.Println(result.RowsAffected)
//	fmt.Println("deleted.")
//
//	json.NewEncoder(w).Encode(book.ID)
//}
