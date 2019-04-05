package routes

import (
	"encoding/json"
	"log"
	"net/http"

	c "github.com/andrelgl/goHelloWord/constants"
	d "github.com/andrelgl/goHelloWord/dao"
	m "github.com/andrelgl/goHelloWord/models"
	"github.com/gorilla/mux"
)

// NewUserRouter create all user routes from API
func NewUserRouter(router *mux.Router) {
	router.HandleFunc("/books", getBooks).Methods("GET")
}
func getBooks(w http.ResponseWriter, r *http.Request) {
	var book m.Book
	books := make([]m.Book, 0)

	if db, ok := d.NewDb(c.PostgreURL); ok {
		rows, err := db.Query("select * from books")

		if err != nil {
			w.WriteHeader(500)
			return
		}
		defer db.Close()

		for rows.Next() {
			if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year); err != nil {
				log.Fatal(err)
			}
			books = append(books, book)
		}
	}

	json.NewEncoder(w).Encode(books)
}
