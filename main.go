package main

import (
	"log"
	"net/http"

	r "github.com/andrelgl/goHelloWord/routes"
	"github.com/gorilla/mux"
)

// ~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~
// package main
// ~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~

func main() {
	router := mux.NewRouter()
	r.NewUserRouter(router)
	log.Fatal(http.ListenAndServe(":8000", router))
}
