package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/todolist/backend/src/model"
)

type tasks []model.Task

var allTask tasks

func retrieveAllTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(allTask)
}

func handleError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotAcceptable)
	fmt.Fprintf(w, "%v", err)
}

func hasError(err error) bool {
	return err != nil
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if hasError(err) {
		handleError(err, w)
	}
	allTask = append(allTask, task)
	json.NewEncoder(w).Encode(allTask)
}

func isValidSearch(id int) bool {
	return len(allTask) > id || allTask != nil
}

// retrieveTask busca uma tarefa pelo seu índice
func retrieveTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if hasError(err) {
		handleError(err, w)
	}
	if !isValidSearch(id) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Not found")
	} else {
		json.NewEncoder(w).Encode(allTask[id])
	}
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.Atoi(params["id"])
	handleError(err, w)

	if isValidSearch(ID) {
		var updated model.Task
		err = json.NewDecoder(r.Body).Decode(&updated)
		handleError(err, w)
		allTask[ID] = updated
		json.NewEncoder(w).Encode(allTask)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not found")
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, _ := strconv.Atoi(params["id"])
	if isValidSearch(ID) {
		allTask = append(allTask[:ID], allTask[ID+1:]...)
	}
	json.NewEncoder(w).Encode(allTask)
}

func main() {
	router := mux.NewRouter()

	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	router.HandleFunc("/tasks/", retrieveAllTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", retrieveTask).Methods("GET")
	router.HandleFunc("/tasks/", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")

	http.ListenAndServe(":8989", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router))
}
