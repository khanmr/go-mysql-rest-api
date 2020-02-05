package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:"id"`
	ISBN   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func dbConn() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:mysqlpassword@/gorestapi")
	if err != nil {
		panic(err.Error())
	}
	return db
}

//Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := dbConn()
	stmt, err := db.Query("SELECT * FROM books ORDER BY id")
	if err != nil {
		panic(err.Error())
	}
	books := []*Book{}
	for stmt.Next() {
		b := &Book{}
		err = stmt.Scan(&b.ID, &b.ISBN, &b.Title, &b.Author)
		if err != nil {
			panic(err.Error())
		}
		books = append(books, b)
	}
	json.NewEncoder(w).Encode(books)
	defer db.Close()
}

//Get Book by ID
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := dbConn()
	params := mux.Vars(r)
	row, err := db.Query("SELECT * FROM books WHERE id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	b := &Book{}
	for row.Next() {
		err = row.Scan(&b.ID, &b.ISBN, &b.Title, &b.Author)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(b)
	defer db.Close()
}

// Add new book
func addBook(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	stmt, err := db.Prepare("INSERT INTO books(isbn, title, author) VALUES(?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	book := make(map[string]string)
	json.Unmarshal(body, &book)
	isbn := book["isbn"]
	title := book["title"]
	author := book["author"]
	_, err = stmt.Exec(isbn, title, author)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	http.Redirect(w, r, "/books", 301)
}

// Update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE books SET isbn = ?, title = ?, author = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	book := make(map[string]string)
	json.Unmarshal(body, &book)
	newISBN := book["isbn"]
	newTitle := book["title"]
	newAuthor := book["author"]
	_, err = stmt.Exec(newISBN, newTitle, newAuthor, params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	http.Redirect(w, r, "/books", 301)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	params := mux.Vars(r)

	stmt, err := db.Prepare("DELETE FROM books WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	http.Redirect(w, r, "/books", 301)
}

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/books", getBooks).Methods("GET")
	mux.HandleFunc("/books/{id}", getBook).Methods("GET")
	mux.HandleFunc("/books", addBook).Methods("POST")
	mux.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	mux.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	log.Println("Starting server on :3000")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
