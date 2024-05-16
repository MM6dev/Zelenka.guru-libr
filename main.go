package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	initializeDatabase()
	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/books", fetchAllBooksHandler).Methods("GET")
	router.HandleFunc("/book", createBookHandler).Methods("POST")
	router.HandleFunc("/book/{id}", fetchBookHandler).Methods("GET")
	router.HandleFunc("/book/{id}", removeBookHandler).Methods("DELETE")
	router.HandleFunc("/genres", fetchAllGenresHandler).Methods("GET")
	router.HandleFunc("/books/genre/{genre}", fetchBooksByGenreHandler).Methods("GET")
	router.HandleFunc("/books/search", searchBooksHandler).Methods("GET")

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func fetchAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := getAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(books)
}

func createBookHandler(w http.ResponseWriter, r *http.Request) {
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := addBook(book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func fetchBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]
	book, err := getBookByID(bookID)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func removeBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]
	if err := deleteBook(bookID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func fetchAllGenresHandler(w http.ResponseWriter, r *http.Request) {
	genres, err := getAllGenres()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(genres)
}

func fetchBooksByGenreHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	genre := vars["genre"]
	books, err := getBooksByGenre(genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(books)
}

func searchBooksHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter 'q' is missing", http.StatusBadRequest)
		return
	}
	books, err := searchBooks(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(books)
}
