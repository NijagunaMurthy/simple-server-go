package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	routes(router)

	http.ListenAndServe(":3050", router)
}
