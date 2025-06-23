package app

import (
	"log"
	"net/http"
	"taskapi/internal/scr/tasks"

	"github.com/gorilla/mux"
)

func Run() {
	router := mux.NewRouter()

	router.HandleFunc("/tasks/{id}", tasks.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", tasks.GetTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", tasks.DeleteTask).Methods("DELETE")
	router.HandleFunc("/tasks", tasks.GetAllTasks).Methods("GET")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
