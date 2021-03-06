package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/personal-projects/go-crud-starter/src/dao"
	"log"
	"net/http"
	"strconv"
)

func HandleGetWorkByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	workID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal("Error parsing id key as int", err)
	}

	work, err := dao.ReadWork(workID)
	if err != nil {
		log.Fatal("Error reading work", err)
	}

	json.NewEncoder(w).Encode(work)
}
