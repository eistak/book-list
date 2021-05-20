package handlers

import (
	"books-list/models"
	"books-list/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Handler struct{}

var books []models.Book

// GetBooks finds all books
func (h Handler) GetBooks(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books = []models.Book{}
		db.Find(&books)
		json.NewEncoder(w).Encode(books)

		log.Println("GetBooks func.")
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, books)
	}
}

// GetBook finds a book by id
func (h Handler) GetBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book

		json.NewDecoder(r.Body).Decode(&book)
		result := db.Create(&book)
		if result.Error != nil {
			log.Println(result.Error)
		}
		log.Println("GetBook func.")
		log.Println(result.RowsAffected)

		json.NewEncoder(w).Encode(book.ID)
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, books)
	}
}

// AddBook adds a book
func (h Handler) AddBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book

		json.NewDecoder(r.Body).Decode(&book)
		result := db.Create(&book)
		if result.Error != nil {
			log.Println(result.Error)
		}
		log.Println("AddBook func.")
		log.Println(result.RowsAffected)

		json.NewEncoder(w).Encode(book.ID)
		w.Header().Set("Content-Type", "text/plan")
		utils.SendSuccess(w, book.ID)
	}
}

// UpdateBook updates a book by id
func (h Handler) UpdateBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book

		json.NewDecoder(r.Body).Decode(&book)
		db.Save(&book)
		log.Println("UpdateBook func.")
		log.Println(book.ID)

		json.NewEncoder(w).Encode(book.ID)
	}
}

// RemoveBook deletes a book by id
func (h Handler) RemoveBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		params := mux.Vars(r)

		var err error
		book.ID, err = strconv.Atoi(params["id"])
		utils.LogFatal(err)

		result := db.Delete(&book, book.ID)
		if result.Error != nil {
			log.Println(result.Error)
		}
		log.Println("RemoveBook func.")
		log.Println(book.ID)
		log.Println(result.RowsAffected)

		json.NewEncoder(w).Encode(book.ID)
		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, book.ID)
	}
}
