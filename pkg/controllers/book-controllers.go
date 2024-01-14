package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ListKelvin/book-store/pkg/models"
	"github.com/ListKelvin/book-store/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book 
func GetBooks(w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks()
	res, _ :=json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0 ,0)
	if err != nil {
		fmt.Println("Error while converting")

	}

	bookDetails, _ := models.GetBookById(ID)
	res, _ :=json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)

}
func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b,_:= CreateBook.CreateBook()
	res, _ :=json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)

}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0 ,0)
	if err != nil {
		fmt.Println("Error while converting")

	}

	book:= models.DeleteBook(ID)
	res, _ :=json.Marshal(book)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0 ,0)
	if err != nil {
		fmt.Println("Error while converting")

	}

	book,db:= models.GetBookById(ID)

	if updateBook.Name != "" {
		book.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		book.Name = updateBook.Author
	}

	if updateBook.Publication != "" {
		book.Name = updateBook.Publication
	}



	db.Save(&book)
	res, _ :=json.Marshal(book)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)

}