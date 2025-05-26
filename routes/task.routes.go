package routes

import (
	"encoding/json"
	"net/http"

	"github.com/NaiG04/go_gorm_restapi/db"
	"github.com/NaiG04/go_gorm_restapi/models"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)

}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task []models.Task
	json.NewEncoder(w).Encode(&task)
	createdTask := db.DB.Create(&task)
	err := createdTask.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func PostTasksHandler(w http.ResponseWriter, r *http.Request) {

}

func DeleteTasksHandler(w http.ResponseWriter, r *http.Request) {

}
