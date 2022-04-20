package routes

import (
	"log"
	"net/http"

	"github.com/bitly/go-simplejson"
)

func Dummy(w http.ResponseWriter, r *http.Request) {
	json := simplejson.New()
	json.Set("status", "ok")

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
	w.WriteHeader(http.StatusOK)
}
