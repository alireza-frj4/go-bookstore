package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/alireza-frj4/go-bookstore/pkg/models"
	"github.com/alireza-frj4/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var Newbook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	Newbooks := models.GetAllBooks()
	res, _ := json.Marshal(Newbooks)
	w.Header().Set("Content-Type", "pkgication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("errir while parsing")
	}
	bookDeatalis, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDeatalis)
	w.Header().Set("Content-Type", "pkgication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Creatbook(w http.ResponseWriter, r *http.Request) {
	Creatbook := &models.Book{}
	utils.ParseBody(r, Creatbook)
	b := Creatbook.Creatbook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("errir while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkgication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var UpdateBook = &models.Book{}
	utils.ParseBody(r, UpdateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("errir while parsing")
	}
	bookDeatalis, db := models.GetBookById(ID)
	if UpdateBook.Name != "" {
		bookDeatalis.Name = UpdateBook.Name
	}
	if UpdateBook.Author != "" {
		bookDeatalis.Author = UpdateBook.Author
	}
	if UpdateBook.Publication != "" {
		bookDeatalis.Publication = UpdateBook.Publication
	}
	db.Save(&bookDeatalis)
	res, _ := json.Marshal(bookDeatalis)
	w.Header().Set("Content-Type", "pkgication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
