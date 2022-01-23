package adminServer

import "net/http"

func PreflightCors(w *http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		(*w).Header().Add("Connection", "keep-alive")
		(*w).Header().Add("Accept", "*/*")
		(*w).Header().Add("Accept-Encoding", "gzip, deflate, br")
		(*w).Header().Add("Access-Control-Allow-Origin", "*")
		(*w).Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET, DELETE, PUT, PATCH")
		(*w).Header().Add("Access-Control-Allow-Headers", "Accept, Connection, Content-Type, Host, Origin, Referer, Sec-Fetch-Dest, Sec-Fetch-Mode, Content-Length, Accept-Encoding, X-CSRF-Token")
		(*w).WriteHeader(http.StatusOK)
	}

	(*w).Header().Add("Access-Control-Allow-Origin", "*")
	(*w).Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET, DELETE, PUT, PATCH")
	(*w).Header().Add("Access-Control-Allow-Headers", "Accept, Connection, Content-Type, Host, Origin, Referer, Sec-Fetch-Dest, Sec-Fetch-Mode, Content-Length, Accept-Encoding, X-CSRF-Token")

	(*w).Header().Set("Content-Type", "application/json")
}
