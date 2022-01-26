package adminServer

import (
	"encoding/json"
	"net/http"

	"garbagelb/config"
)

func GetListeners(w http.ResponseWriter, r *http.Request) {

	listeners := config.GetListeners()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(listeners)
}
