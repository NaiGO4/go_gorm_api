package main

import (
	"net/http"

	"github.com/NaiG04/go_gorm_restapi/db"
	"github.com/NaiG04/go_gorm_restapi/models"
	"github.com/NaiG04/go_gorm_restapi/routes"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database connection
	db.DBConnection()
	// Conexion a los modelos
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()
	// Rutas de la API
	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/user/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/user/{id}", routes.DeleteUserHandler).Methods("DELETE")

	// Rutas de las tareas
	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/task/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.PostTasksHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.DeleteTasksHandler).Methods("DELETE")

	http.ListenAndServe(":3030", r)
}
