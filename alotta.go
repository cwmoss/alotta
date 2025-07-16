package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {
	var port = "9999"

	fmt.Printf("starting server on http://localhost:%s\n", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User
		json.NewDecoder(r.Body).Decode(&user)

		fmt.Fprintf(w, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
	})

	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		peter := User{
			Firstname: "John",
			Lastname:  "Doe",
			Age:       25,
		}

		json.NewEncoder(w).Encode(peter)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":"+port, nil)

	/*
		router := routegroup.New(http.NewServeMux())
		router.Use(loggingMiddleware)

		// handle the /hello route
		router.Handle("GET /hello", helloHandler)

		// create a new group for the /api path
		apiRouter := router.Mount("/api")
		// add middleware
		apiRouter.Use(loggingMiddleware, corsMiddleware)

		// route handling
		apiRouter.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, API!"))
		})

		// add another group with its own set of middlewares
		protectedGroup := router.Group()
		protectedGroup.Use(authMiddleware)
		protectedGroup.HandleFunc("GET /protected", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Protected API!"))
		})

		http.ListenAndServe(":8080", router)
	*/
}
