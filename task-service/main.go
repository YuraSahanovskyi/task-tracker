package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
)

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

var tasks []Task = []Task{
	{ID: "1", Title: "Task 1", Description: "Description for task 1", Status: "pending"},
	{ID: "2", Title: "Task 2", Description: "Description for task 2", Status: "completed"},
}

func getAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(tasks)
	if err != nil {
		w.WriteHeader(400)
	}
}

func getTaskById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")
	for _, task := range tasks {
		if task.ID == id {
			err := json.NewEncoder(w).Encode(task)
			if err != nil {
				w.WriteHeader(400)
			}
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newTask Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Error decoding request body:", err)
		return
	}

	newTask.ID = fmt.Sprintf("%d", len(tasks)+1)
	tasks = append(tasks, newTask)

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(newTask)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error encoding JSON:", err)
	}
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")

	var updatedTask Task
	err := json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Error decoding request body:", err)
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			updatedTask.ID = id
			tasks[i] = updatedTask
		}
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(updatedTask)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error encoding JSON:", err)
	}
}
func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")
	for i, task := range tasks {
		if id == task.ID {
			tasks = slices.Delete(tasks, i, i+1)
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	fmt.Println("Hello from task service")
	mux := http.NewServeMux()
	mux.HandleFunc("GET /tasks", getAllTasks)
	mux.HandleFunc("GET /tasks/{id}", getTaskById)
	mux.HandleFunc("POST /tasks", createTask)
	mux.HandleFunc("PUT /tasks/{id}", updateTask)
	mux.HandleFunc("DELETE /tasks/{id}", deleteTask)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
