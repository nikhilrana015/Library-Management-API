package routes

import (
	"github.com/gorilla/mux"
	controllers "github.com/nikhilrana/Go-LibraryMgmt/pkg/controllers"
)

var RegisterBookInLibraryRoutes = func(router *mux.Router) {
	router.HandleFunc("/books/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/", controllers.AddNewBook).Methods("POST")
	router.HandleFunc("/books/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBookById).Methods("DELETE")
	router.HandleFunc("/checkin", controllers.CheckInBook).Methods("PATCH")
	router.HandleFunc("/checkout", controllers.CheckOutBook).Methods("PATCH")
}
