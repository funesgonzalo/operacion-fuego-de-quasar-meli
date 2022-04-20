package utils

import "net/http"

func HeaderConfig(res http.ResponseWriter) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(res).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	res.Header().Set("Content-Type", "application/json")
}
