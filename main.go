package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"github.com/realr3fo/tkai_circles_account/app"
	"github.com/realr3fo/tkai_circles_account/controllers"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", controllers.HelloWorld).Methods("GET")
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":" + port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
