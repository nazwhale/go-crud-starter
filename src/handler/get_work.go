package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/personal-projects/go-crud-starter/src/dao"
	"log"
	"net/http"
	"strconv"
)

func HandleGetWork(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	workID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal("Error parsing id key as int", err)
	}

	work, err := dao.ReadWork(workID)
	if err != nil {
		log.Fatal("Error listing works", err)
	}

	json.NewEncoder(writer).Encode(work)
}
