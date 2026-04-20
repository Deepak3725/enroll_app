package routes

import (
	"fmt"
	"log"
	controller "myapp/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes() {
	port := "8080"
	router := mux.NewRouter()

	// Register route
	router.HandleFunc("/student/add", controller.AddStudent).Methods("POST")
	router.HandleFunc("/student/{sid}", controller.GetStud).Methods("GET")
	router.HandleFunc("/students", controller.GetAllStuds)

	router.HandleFunc("/course/add", controller.AddCourse).Methods("POST")
	router.HandleFunc("/course/{sid}", controller.GetCourse).Methods("GET")

	// Update route
	router.HandleFunc("/student/{sid}", controller.UpdateStud).Methods("PUT")

	router.HandleFunc("/student/{sid}", controller.DeleteStud).Methods("DELETE")


	//Serving static files
	fhandler:=http.FileServer(http.Dir("./views"))
	router.PathPrefix("/").Handler(fhandler)
	log.Println("Application running on port " + port)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
