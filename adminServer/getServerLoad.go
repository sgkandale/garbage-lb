package adminServer

import (
	"encoding/json"
	"net/http"
	"strings"

	"garbagelb/serverLoad"
)

func GetServerLoad(w http.ResponseWriter, r *http.Request) {

	if strings.ToLower(r.Method) == "get" || strings.ToLower(r.Method) == "options" {

		load, err := serverLoad.GetServerLoad()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(load)

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
