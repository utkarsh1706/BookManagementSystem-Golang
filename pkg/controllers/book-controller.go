package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/utkarsh1706/Golang-BookManagementSystem/pkg/models"
	"github.com/utkarsh1706/Golang-BookManagementSystem/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	books, err := models.GetAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	bookDetails, err := models.GetBookById(bookId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var createBook models.Book
	utils.ParseBody(r, &createBook)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := models.CreateBook(ctx, &createBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	createBook.ID = result.InsertedID.(primitive.ObjectID)
	res, _ := json.Marshal(createBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := models.DeleteBook(ctx, bookId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)
	bookDetails, err := models.GetBookById(bookId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	err = models.UpdateBook(ctx, bookId, bookDetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
