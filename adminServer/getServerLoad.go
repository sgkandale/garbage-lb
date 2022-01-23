package adminServer

import (
	"encoding/json"
	"net/http"

	"garbagelb/serverLoad"
)

func GetServerLoad(w http.ResponseWriter, r *http.Request) {

	// development purpose
	w.Header().Add("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	load, err := serverLoad.GetServerLoad()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(load)
}
