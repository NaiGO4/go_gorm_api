package main

import (
	"net/http"

	"github.com/NaiG04/go_gorm_restapi/db"
	"github.com/NaiG04/go_gorm_restapi/models"
	"github.com/NaiG04/go_gorm_restapi/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3030", r)
}
