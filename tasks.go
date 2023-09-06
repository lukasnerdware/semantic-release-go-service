package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Task represents a single todo task.
type Task struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var tasks []Task

func main() {
	log.Println("Starting the Todo App...")
	r := mux.NewRouter()

	// Define API routes
	r.HandleFunc("/tasks", GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", GetTask).Methods("GET")
	r.HandleFunc("/tasks", CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", DeleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}

// GetTasks returns a list of all tasks.
func GetTasks(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting all tasks...")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// GetTask returns a single task by its ID.
func GetTask(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range tasks {
		if item.ID == params["id"] {
			json.NewEncoder(response).Encode(item)
			return
		}
	}
	json.NewEncoder(response).Encode(&Task{})
}

// CreateTask creates a new task.
func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = fmt.Sprintf("%d", len(tasks)+1)
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(task)
}

// UpdateTask updates an existing task by its ID.
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range tasks {
		if item.ID == params["id"] {
			var updatedTask Task
			_ = json.NewDecoder(r.Body).Decode(&updatedTask)
			updatedTask.ID = params["id"]
			tasks[index] = updatedTask
			json.NewEncoder(w).Encode(updatedTask)
			return
		}
	}
	json.NewEncoder(w).Encode(&Task{})
}

// DeleteTask deletes a task by its ID.
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range tasks {
		if item.ID == params["id"] {
			tasks = append(tasks[:index], tasks[index+1:]...)
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
}
