package routes

import (
	"fmt"
	"log"
	controller "myapp/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes() {
	port := "8080"            //assigning port number in this port we will start our code
	router := mux.NewRouter() //gorrila mux creates a router that will take different incoming request

	// Register route
	router.HandleFunc("/student/add", controller.AddStudent).Methods("POST") //url,fx return but only during Method (post)
	router.HandleFunc("/student/{sid}", controller.GetStud).Methods("GET")
	router.HandleFunc("/students", controller.GetAllStuds) //No restriction in every method it returns

	router.HandleFunc("/course/add", controller.AddCourse).Methods("POST")
	router.HandleFunc("/course/{sid}", controller.GetCourse).Methods("GET")

	// Update route
	router.HandleFunc("/student/{sid}", controller.UpdateStud).Methods("PUT") //gorrila mux provide us with different methods and we can use this sid in other cases in controller which server mux don't provide
	router.HandleFunc("/course/{sid}", controller.UpdateCourse).Methods("UPDATE")

	router.HandleFunc("/course/{sid}", controller.DeleteCourse).Methods("DELETE")
	router.HandleFunc("/signup", controller.Signup).Methods("POST")

	router.HandleFunc("/login", controller.Login).Methods("POST")

	router.HandleFunc("/student/{sid}", controller.DeleteStud).Methods("DELETE")

	router.HandleFunc("/courses", controller.GetAllCourses)
	router.HandleFunc("/logout",controller.Logout)

	router.HandleFunc("/enroll/{sid}/{cid}",controller.GetEnroll).Methods("GET")
	router.HandleFunc("/enrolls",controller.GetEnrolls).Methods("GET")
	router.HandleFunc("/enroll/{sid}/{cid}",controller.DeleteEnroll).Methods("DELETE")


	//Serving static files
	fhandler := http.FileServer(http.Dir("./view"))
	router.PathPrefix("/").Handler(fhandler)
	log.Println("Application running on port " + port)

	err := http.ListenAndServe(":8080", router) // port ,use the mux of router if nil uses default mux
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
