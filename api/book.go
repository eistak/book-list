package api

import (
	"books-list/models"
	"books-list/tools"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Api struct{}

var books []models.Book

// GetBooks finds all books
func (api Api) GetBooks(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books = []models.Book{}
		db.Find(&books)
		er := json.NewEncoder(w).Encode(books)
		tools.LogFatal(er)

		log.Println("GetBooks func.")
		w.Header().Set("Content-Type", "application/json")
	}
}

// GetBook finds a book by id
func (api Api) GetBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var err models.Error

		er := json.NewDecoder(r.Body).Decode(&book)
		tools.LogFatal(er)

		result := db.Create(&book)
		if result.Error != nil {
			log.Println(result.Error)
			tools.SendError(w, http.StatusInternalServerError, err)
			return
		}
		log.Println("GetBook func.")
		log.Println(result.RowsAffected)

		er = json.NewEncoder(w).Encode(book.ID)
		tools.LogFatal(er)

		w.Header().Set("Content-Type", "application/json")
	}
}

// AddBook adds a book
func (api Api) AddBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var err models.Error

		er := json.NewDecoder(r.Body).Decode(&book)
		tools.LogFatal(er)

		if book.Author == "" || book.Title == "" || book.Year == "" {
			err.Message = "All fileds are required."
			tools.SendError(w, http.StatusBadRequest, err)
			return
		}

		result := db.Create(&book)
		if result.Error != nil {
			log.Println(result.Error)
			tools.SendError(w, http.StatusInternalServerError, err)
			return
		}
		log.Println("AddBook func.")
		log.Println(result.RowsAffected)

		er = json.NewEncoder(w).Encode(book.ID)
		tools.LogFatal(er)

		w.Header().Set("Content-Type", "text/plan")
	}
}

// UpdateBook updates a book by id
func (api Api) UpdateBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var err models.Error

		er := json.NewDecoder(r.Body).Decode(&book)
		tools.LogFatal(er)

		if book.ID == 0 || book.Author == "" || book.Title == "" || book.Year == "" {
			err.Message = "All fileds are required."
			tools.SendError(w, http.StatusBadRequest, err)
			return
		}

		db.Save(&book)
		log.Println("UpdateBook func.")
		log.Println(book.ID)

		er = json.NewEncoder(w).Encode(book.ID)
		tools.LogFatal(er)
	}
}

// RemoveBook deletes a book by id
func (api Api) RemoveBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var err models.Error
		params := mux.Vars(r)

		book.ID, _ = strconv.Atoi(params["id"])

		result := db.Delete(&book, book.ID)
		if result.Error != nil {
			log.Println(result.Error)
			tools.SendError(w, http.StatusInternalServerError, err)
			return
		}
		log.Println("RemoveBook func.")
		log.Println(book.ID)
		log.Println(result.RowsAffected)

		er := json.NewEncoder(w).Encode(book.ID)
		tools.LogFatal(er)

		w.Header().Set("Content-Type", "text/plain")
	}
}
