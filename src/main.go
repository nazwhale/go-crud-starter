package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/personal-projects/go-crud-starter/src/handler"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Connect()
	if err != nil {
		// handle
	}
	dao := dao.New(db)
	handler := handler.New(dao)
	router.HandleFunc("/works/list", handler.HandleList)
	// consider passing up errors and handling them here
	handleRequests()
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	// consider using mux subrouters for seperate art and artists
	router.HandleFunc("/", handler.HandleRoot)
	router.HandleFunc("/works/list", handler.HandleListWorks)
	router.HandleFunc("/work/create", handler.HandleCreateWork).Methods("POST")
	router.HandleFunc("/work/update", handler.HandleUpdateWork).Methods("POST")
	router.HandleFunc("/work/{id}", handler.HandleGetWorkByID)

	log.Fatal("ListenAndServe error: ", http.ListenAndServe(getPort(), router))
}

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "4747"
		fmt.Println("Defaulting to http://localhost:" + port)
	}
	return ":" + port
}

// notes
// consider removing src level
// see this for structuring best practices https://peter.bourgon.org/go-for-industrial-programming/
//
// login flow + auth
// password recovery by integrating with email client
// handling the client going down or timing out

package dao

type DAO interface {
	ReadWork(id int) (*Work, error)
}

func New(db *sql.DB) DAO {
	return &dao{db}
}

type dao struct {
	db *sql.DB
}

func (d dao) ReadWork(id int) (*Work, error) {
	d.db.QueryRow()
}


package handler

type Handler struct {
	dao dao.DAO
}

func New(dao dao.DAO) Handler {
	return &Handler{dao}
}

func (h *Handler) HandleList(w http.ResponseWriter, r* http.Request) {}
h.dao.List
}
