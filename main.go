package main

import (
	"net/http"

	"github.com/francixcode2004/go-api-rest/db"
	"github.com/francixcode2004/go-api-rest/models"
	"github.com/francixcode2004/go-api-rest/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()
	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Task{})
	router := mux.NewRouter()
	router.HandleFunc("/", routes.HomeHandler)

	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/", routes.PostUserHandler).Methods("POST")
	router.HandleFunc("/", routes.DeleteUserHandler).Methods("DELETE")
	http.ListenAndServe(":8080", router)

}
