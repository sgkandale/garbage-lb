package adminServer

import (
	"encoding/json"
	"net/http"

	"garbagelb/config"
)

func GetListeners(w http.ResponseWriter, r *http.Request) {

	// development purpose
	w.Header().Add("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	listeners := config.GetListeners()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(listeners)
}
