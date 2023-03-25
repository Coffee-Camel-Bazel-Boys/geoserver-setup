package main

import(
	"net/http"
	"mapController"

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/geo", mapController.HandleRequest);

	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":6001", handler);

}