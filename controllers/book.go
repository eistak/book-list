package controllers

import (
	"books-list/models"
	"books-list/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Controller struct{}

var books []models.Book

// GetBooks finds all books
func (c Controller) GetBooks(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books = []models.Book{}
		db.Find(&books)

		json.NewEncoder(w).Encode(books)
	}
}

// GetBook find a book by id
func (c Controller) GetBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book

		json.NewDecoder(r.Body).Decode(&book)
		result := db.Create(&book)
		if result.Error != nil {
			log.Println(result.Error)
		}
		log.Println(result.RowsAffected)
		fmt.Println("added.")

		json.NewEncoder(w).Encode(book.ID)
	}
}

// AddBook adds a book
func (c Controller) AddBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book

		json.NewDecoder(r.Body).Decode(&book)
		result := db.Create(&book)
		if result.Error != nil {
			log.Println(result.Error)
		}
		log.Println(result.RowsAffected)
		fmt.Println("added.")

		json.NewEncoder(w).Encode(book.ID)
	}
}

// UpdateBook updates a book by id
func (c Controller) UpdateBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book

		json.NewDecoder(r.Body).Decode(&book)
		db.Save(&book)

		fmt.Println(book.ID)
		log.Println("updated.")
		json.NewEncoder(w).Encode(book.ID)
	}
}

// RemoveBook deletes a book by id
func (c Controller) RemoveBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		params := mux.Vars(r)

		var err error
		book.ID, err = strconv.Atoi(params["id"])
		utils.LogFatal(err)

		log.Println(book.ID)
		result := db.Delete(&book, book.ID)
		if result.Error != nil {
			log.Println(result.Error)
		}
		log.Println(result.RowsAffected)
		fmt.Println("deleted.")

		json.NewEncoder(w).Encode(book.ID)
	}
}
