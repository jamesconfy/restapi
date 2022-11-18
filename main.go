package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Author struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Book struct {
	ID     int    `json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author *Author
}

var Books []Book

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// fmt.Fprint(w, str)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Hello, This is the homepage, please ignore")
	}
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		if Books == nil {
			json.NewEncoder(w).Encode("[]")
			return
		}
		json.NewEncoder(w).Encode(&Books)
	}

	if r.Method == "POST" {
		var book Book
		_ = json.NewDecoder(r.Body).Decode(&book)
		book.ID = rand.Intn(1000000000)
		Books = append(Books, book)
		json.NewEncoder(w).Encode(&book)
	}
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if r.Method == "GET" {
		for _, item := range Books {
			num, _ := strconv.Atoi(params["id"])
			if item.ID == num {
				json.NewEncoder(w).Encode(&item)
				return
			}
		}
	}

	if r.Method == "DELETE" {
		for index, item := range Books {
			num, _ := strconv.Atoi(params["id"])
			if item.ID == num {
				Books = append(Books[:index], Books[index+1:]...)
				break
			}
		}

		json.NewEncoder(w).Encode(&Books)
	}

}

func main() {
	router := mux.NewRouter()
	Books = append(Books, Book{ID: 1, Isbn: "1234", Title: "This is the first book", Author: &Author{ID: "1", Name: "Confidence James"}})
	Books = append(Books, Book{ID: 2, Isbn: "2234", Title: "This is the second book", Author: &Author{ID: "1", Name: "Confidence James"}})
	Books = append(Books, Book{ID: 3, Isbn: "3334", Title: "This is the third book", Author: &Author{ID: "1", Name: "Confidence James"}})
	Books = append(Books, Book{ID: 4, Isbn: "4234", Title: "This is the fourth book", Author: &Author{ID: "1", Name: "Confidence James"}})
	Books = append(Books, Book{ID: 5, Isbn: "5234", Title: "This is the fifth book", Author: &Author{ID: "1", Name: "Confidence James"}})

	router.HandleFunc("/api/", home).Methods("GET")
	router.HandleFunc("/api/books", getBooks).Methods("GET", "POST")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET", "PATCH", "DELETE")
	http.Handle("/api/", router)
	// http.Handle("/api/books", router)
	http.ListenAndServe(":8080", nil)
}
