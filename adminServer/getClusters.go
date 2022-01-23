package adminServer

import (
	"encoding/json"
	"net/http"

	"garbagelb/config"
)

func GetClusters(w http.ResponseWriter, r *http.Request) {

	clusters := config.GetClusters()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clusters)
}
