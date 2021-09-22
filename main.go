package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Struct (Model)

type Book struct {
	Id string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

// Init books variable as a slice of Book struct
var books []Book // Used inside for range iteration

var x, y string = "Content-Type", "application/json"

func getAllBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set(x, y)
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set(x, y)
	params := mux.Vars(r) // Get params
	for i, item := range books {
		if params["id"] == item.Id {
			json.NewEncoder(w).Encode(books[i])
			return
		} 
	}
	// json.NewEncoder(w).Encode(&Book{}) // without this specific line, the code is also running
}

func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set(x, y)
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	book.Id = strconv.Itoa(len(books) + 1)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set(x, y)
	params := mux.Vars(r)
	for i, item := range books{
		if params["id"] == item.Id{
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	book.Id = params["id"]
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set(x, y)
	params := mux.Vars(r)
	for i, item := range books{
		if params["id"] == item.Id{
			books = append(books[:i], books[i+1:]...)
			json.NewEncoder(w).Encode(books)
			return
		}
	}
	// json.NewEncoder(w).Encode(books)
}

func main(){
	// Init Router
	r:=mux.NewRouter()

	// Mock data @todo - implement DB
	//Book 1
	books = append(books, Book{
		Id: "1", Isbn: "55418", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	
	//Book 2
	books = append(books, Book{
		Id: "2", Isbn: "58135", Title: "Book Two", Author: &Author{Firstname: "Lisa", Lastname: "Kudrow"}})
	
	//Book 3
	books = append(books, Book{
		Id: "3", Isbn: "51574", Title: "Book Three", Author: &Author{Firstname: "Lisa", Lastname: "Simpson"}})


	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getAllBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/book/crt", createBook).Methods("POST")
	r.HandleFunc("/api/book/upd/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/book/del/{id}", deleteBook).Methods("DELETE")

	var port string = ":1928"
	// fmt.Printf("Server running on port%v\n", port) //Must be removed during production
	log.Fatal(http.ListenAndServe(port, r))
}