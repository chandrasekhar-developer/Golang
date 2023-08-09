package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/chandrasekhar-developer/go-bookstore-mysql/pkg/models"
	"github.com/chandrasekhar-developer/go-bookstore-mysql/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing...")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.ParseInt(params["id"], 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	UpdateBook := &models.Book{}
	utils.ParseBody(r, UpdateBook)

	params := mux.Vars(r)

	ID, err := strconv.ParseInt(params["id"], 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")
	}

	book, db := models.GetBookById(ID)

	if UpdateBook.Name != "" {
		book.Name = UpdateBook.Name
	}

	if UpdateBook.Author != "" {
		book.Author = UpdateBook.Author
	}

	if UpdateBook.Publication != "" {
		book.Publication = UpdateBook.Publication
	}

	db.Save(&book)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
