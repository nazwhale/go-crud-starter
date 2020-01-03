package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/personal-projects/go-crud-starter/src/handler"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	handleRequests()
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handler.HandleRoot)
	router.HandleFunc("/works/list", handler.HandleListWorks)
	router.HandleFunc("/work/create", handler.HandleCreateWork).Methods("POST")
	router.HandleFunc("/work/{id}", handler.HandleGetWorkByID)

	err := http.ListenAndServe(getPort(), router)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "4747"
		fmt.Println("Defaulting to http://localhost:" + port)
	}
	return ":" + port
}
