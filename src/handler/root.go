package handler

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "👋")
}