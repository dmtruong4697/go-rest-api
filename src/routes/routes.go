package routes

import (
	"database/sql"
	"go-rest-api/src/controllers"

	"github.com/gorilla/mux"
)

func Setup(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	userController := controllers.NewUserController(db)

	router.HandleFunc("/users", userController.CreateUser).Methods("POST")
	router.HandleFunc("/update-users", userController.UpdateUser).Methods("POST")
	router.HandleFunc("/get-user", userController.GetUserByID).Methods("POST")
	router.HandleFunc("/delete-user", userController.DeleteUserByID).Methods("POST")

	return router
}
