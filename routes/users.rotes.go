package routes

import (
	"encoding/json"
	"net/http"

	"github.com/francixcode2004/go-api-rest/db"
	"github.com/francixcode2004/go-api-rest/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
	w.Write([]byte("Get Users Handler"))
}
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r) // get
	db.DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(&user)
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User  not found!"))
		return
	}
	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)
	json.NewEncoder(w).Encode(&user)
	w.Write([]byte("Get User Handler"))
}
func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	createuser := db.DB.Create(&user)
	err := createuser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&user)

	w.Write([]byte("Post Users Handler"))
}
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User  not found!"))
		return
	}

	db.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Delete Users Handler"))
}
