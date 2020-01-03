package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/personal-projects/go-crud-starter/src/dao"
)

func HandleListWorks(w http.ResponseWriter, r *http.Request) {
	works, err := dao.ListWorks(10)

	if err != nil {
		log.Fatal("Error listing works", err)
	}

	json.NewEncoder(w).Encode(works)
}
