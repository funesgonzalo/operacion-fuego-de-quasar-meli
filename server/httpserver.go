package server

import (
	"log"
	"net/http"
	"operacionFuegoQuasar/server/routes"
	topsecrect_post "operacionFuegoQuasar/server/routes/topsecret"
	topsecrect_split "operacionFuegoQuasar/server/routes/topsecret_split"
	"os"

	"github.com/gorilla/mux"
)

const defaultPort = "8080"

func Start() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := mux.NewRouter()
	r.HandleFunc("/", routes.Dummy).Methods("GET")
	r.HandleFunc("/topsecret", topsecrect_post.TopSecretPost).Methods("POST")
	r.HandleFunc("/topsecret_split/{satellite_name}", topsecrect_split.TopSecretSplit_Post).Methods("POST")
	r.HandleFunc("/topsecret_split", topsecrect_split.TopSecretSplit_Get).Methods("GET")
	http.Handle("/", r)

	// Iniciamos el servidor
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
