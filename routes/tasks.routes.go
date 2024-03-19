package routes

import (
	"encoding/json"
	"net/http"

	"github.com/francixcode2004/go-api-rest/db"
	"github.com/francixcode2004/go-api-rest/models"
	"github.com/gorilla/mux"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	cratetask := db.DB.Create(&task)
	err := cratetask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&task)
}
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])

	if task.UserID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User  not found!"))
		return
	}

	json.NewEncoder(w).Encode(&task)
}
func DeleteTasksHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])
	if task.UserID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User  not found!"))
		return
	}
	db.DB.Unscoped().Delete(&task, params["id"])
	w.WriteHeader(http.StatusNoContent)

}
