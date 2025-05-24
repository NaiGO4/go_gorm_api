package main

import (
	"net/http"

	"github.com/NaiG04/go_gorm_restapi/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3030", r)
}
