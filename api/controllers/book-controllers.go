package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ListKelvin/book-store/api/models"
	"github.com/ListKelvin/book-store/api/utils"
	"github.com/gorilla/mux"
)


var NewBook models.Book 
var server = Server{}
func (server *Server) GetBooks(w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks(server.DB)
	res, _ :=json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)
}

func(server *Server) GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0 ,0)
	if err != nil {
		fmt.Println("Error while converting")

	}

	bookDetails, _ := models.GetBookById(ID, server.DB)
	res, _ :=json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)

}
func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b,_:= CreateBook.CreateBook(server.DB)
	res, _ :=json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func(server *Server) DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0 ,0)
	if err != nil {
		fmt.Println("Error while converting")

	}

	book:= models.DeleteBook(ID , server.DB)
	res, _ :=json.Marshal(book)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)

}

// need to fix
func(server *Server) UpdateBook(w http.ResponseWriter, r *http.Request){
	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0 ,0)
	if err != nil {
		fmt.Println("Error while converting")

	}

	book , db:= models.GetBookById(ID,server.DB)

	// if updateBook.Name != "" {
	// 	book.Name = updateBook.Name
	// }
	// if updateBook.Author != "" {
	// 	book.Name = updateBook.Author
	// }

	// if updateBook.Publication != "" {
	// 	book.Name = updateBook.Publication
	// }



	db.Save(&book)
	res, _ :=json.Marshal(book)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)

}