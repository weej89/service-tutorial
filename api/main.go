package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"service-tutorial/api/internal/user"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user", user.CreateUserAccount).Methods("POST")
	http.ListenAndServe(":8080", router)
}
