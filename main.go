package main

import (
	"encoding/json";
	"log";
	"net/http";
	"math/rand";
	"strconv";
	"github.com/gorilla/mux"
)

 // Book Struct (MODEL)
 type Book struct {
	 ID string  `json:"id"`;
	 Isbn string `json:"isbn"`;
	 Title string `json:"title"`;
	 Author *Author `json:"author"`
 }

 // Author struct
 type Author struct {
	 FirstName string `json:"firstname"`
	 LastName string `json:"lastname"`
 }
 
 // inits books var as a slice Book struct
 var books []Book

 // get all books
 func getBooks (w http.ResponseWriter, r *http.Request) {
	 w.Header().Set("Content-Type", "application/json")
	 json.NewEncoder(w).Encode(books)
 }
 
 // get single
 func getBook (w http.ResponseWriter, r *http.Request) {
	 w.Header().Set("Content-Type", "application/json")
	 params := mux.Vars(r) // Get params 
	 // loop through books and find with id
		 for _, item := range books {
			 if item.ID  == params["id"] {
				json.NewEncoder(w).Encode(item)	
				return 			
			 }
		}

		 json.NewEncoder(w).Encode(&Book{})

}

// create book
func createBook (w http.ResponseWriter, r *http.Request) {
	 w.Header().Set("Content-Type", "application/json")
	 var book Book 
	 _ = json.NewDecoder(r.Body).Decode(&book)
	 book.ID = strconv.Itoa(rand.Intn(10000000)) // mock id - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// update book
func updateBooks (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = strconv.Itoa(rand.Intn(10000000)) // mock id - not safe
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
		}
	}
	json.NewEncoder(w).Encode(books)
}

// delete book
func deleteBooks (w http.ResponseWriter, r *http.Request) {
	 w.Header().Set("Content-Type", "application/json")
	 params := mux.Vars(r)
	 for index, item := range books {
		 if item.ID == params["id"] {
			 books = append(books[:index], books[index+1:]...)
			 break
		 }
	 }
	 json.NewEncoder(w).Encode(books)
}

func main() {
	// init Router
	r := mux.NewRouter()

	// mock data = @todo = implements database
	books = append(books, Book{ ID: "1", Isbn: "448793", Title: "Book one!", Author: &Author{ FirstName: "Jonh", LastName: "Doe" } })
	books = append(books, Book{ ID: "2", Isbn: "556090", Title: "Book two!", Author: &Author{ FirstName: "Jonh", LastName: "Doe" } })
	books = append(books, Book{ ID: "3", Isbn: "8722480", Title: "Book three!", Author: &Author{ FirstName: "Jonh", LastName: "Doe" } })

	// var age int = 35
	// age := 35

	// Route Handler / Enpoints
	r.HandleFunc("api/books", getBooks).Methods("GET")
	r.HandleFunc("api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("api/books", createBook).Methods("POST")
	r.HandleFunc("api/books/{id}", updateBooks).Methods("PUT")
	r.HandleFunc("api/books/{id}", deleteBooks).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}