package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	models "github.com/nikhilrana/Go-LibraryMgmt/pkg/models"
	utils "github.com/nikhilrana/Go-LibraryMgmt/pkg/utils"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, req *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["id"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(id)

	var res []byte

	w.Header().Set("Content-Type", "pkglication/json")

	if bookDetails.ID == 0 {
		error := make(map[string]string)
		error["message"] = "Not Found"
		res, _ = json.Marshal(error)
		w.WriteHeader(http.StatusNotFound)

	} else {
		res, _ = json.Marshal(bookDetails)
		w.WriteHeader(http.StatusOK)
	}

	w.Write(res)
}

func AddNewBook(w http.ResponseWriter, req *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(req, createBook)
	b := createBook.AddNewBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteBookById(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	bookId := vars["id"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, _ := models.GetBookById(id)
	var res []byte

	w.Header().Set("Content-Type", "pkglication/json")

	if bookDetails.ID == 0 {
		error := make(map[string]string)
		error["message"] = "Not Found"
		res, _ = json.Marshal(error)
		w.WriteHeader(http.StatusNotFound)

	} else {
		models.DeleteBookById(id)
		res, _ = json.Marshal(bookDetails)
		w.WriteHeader(http.StatusOK)
	}

	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, req *http.Request) {
	updatedBook := &models.Book{}
	utils.ParseBody(req, updatedBook)
	vars := mux.Vars(req)
	bookId := vars["id"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db := models.GetBookById(id)

	if bookDetails.ID == 0 {
		error := make(map[string]string)
		error["message"] = "Not Found"
		res, _ := json.Marshal(error)
		w.Header().Set("content-type", "pkglication/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(res)
		return
	}

	if updatedBook.Name != "" {
		bookDetails.Name = updatedBook.Name
	}
	if updatedBook.Author != "" {
		bookDetails.Author = updatedBook.Author
	}
	if updatedBook.Publication != "" {
		bookDetails.Publication = updatedBook.Publication
	}
	if updatedBook.Quantity != 0 {
		bookDetails.Quantity = updatedBook.Quantity
	}

	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CheckInBook(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	bookId, present := query["id"]

	// fmt.Println(bookId)
	// fmt.Println(present)

	if !present || len(bookId) == 0 {
		w.Header().Set("Content-Type", "pkglication/json")
		error := make(map[string]string)
		error["message"] = "Query Parameter not found"
		res, _ := json.Marshal(error)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}

	id, err := strconv.ParseInt(bookId[0], 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, db := models.GetBookById(id)
	var res []byte

	w.Header().Set("Content-Type", "pkglication/json")

	if bookDetails.ID == 0 {
		error := make(map[string]string)
		error["message"] = "Not Found"
		res, _ = json.Marshal(error)
		w.WriteHeader(http.StatusNotFound)
	} else {
		bookDetails.Quantity += 1
		db.Save(&bookDetails)
		res, _ = json.Marshal(bookDetails)
		w.WriteHeader(http.StatusOK)
	}
	w.Write(res)
}

func CheckOutBook(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	bookId, present := query["id"]

	if !present || len(bookId) == 0 {
		w.Header().Set("Content-Type", "pkglication/json")
		error := make(map[string]string)
		error["message"] = "Query Parameter not found"
		res, _ := json.Marshal(error)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}

	id, err := strconv.ParseInt(bookId[0], 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, db := models.GetBookById(id)
	var res []byte
	error := make(map[string]string)

	w.Header().Set("Content-Type", "pkglication/json")

	if bookDetails.ID == 0 {
		error["message"] = "Not Found"
		res, _ = json.Marshal(error)
		w.WriteHeader(http.StatusNotFound)
	} else if bookDetails.Quantity == 0 {
		error["message"] = "Book not available"
		res, _ = json.Marshal(error)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		bookDetails.Quantity -= 1
		db.Save(&bookDetails)
		res, _ = json.Marshal(bookDetails)
		w.WriteHeader(http.StatusOK)
	}

	w.Write(res)

}
