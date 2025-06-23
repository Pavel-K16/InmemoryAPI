package tasks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	e "taskapi/internal/entities"

	"github.com/gorilla/mux"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	log.Tracef("Handler: CreateTask")

	params := mux.Vars(r)
	id := params["id"]
	id = strings.Trim(id, "{}")
	if id == "" {
		log.Errorf("Handler: CreateTask: %v", fmt.Errorf("invalid id"))
		http.Error(w, "invalid id", http.StatusBadRequest)

		return
	}

	info := &e.EntityInfo{
		ID: id,
	}

	task, err := Create(info)
	if err != nil {
		log.Warningf("Handler: CreateTask: %v", err)
		http.Error(w, err.Error(), http.StatusConflict)

		return
	}

	taskStatus, ok := task.(*e.TaskStatus)
	if !ok {
		err := fmt.Errorf("invalid data")
		log.Errorf("Handler: CreateTask: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	data, err := json.Marshal(taskStatus)
	if err != nil {
		log.Errorf("Handler: CreateTask: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	log.Tracef("Handler: GetTask")

	params := mux.Vars(r)
	id := params["id"]
	id = strings.Trim(id, "{}")

	if id == "" {
		log.Errorf("Handler: GetTask: %v", fmt.Errorf("invalid id"))
		http.Error(w, "invalid id", http.StatusBadRequest)

		return
	}

	info := &e.EntityInfo{
		ID: id,
	}

	task, err := Get(info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)

		return
	}

	taskStatus, ok := (*task).(e.TaskStatus)
	if !ok {
		err := fmt.Errorf("invalid data")
		log.Errorf("Handler: GetTask: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	data, err := json.Marshal(taskStatus)
	if err != nil {
		log.Errorf("Handler: GetTask: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	log.Tracef("Handler: DeleteTask")

	params := mux.Vars(r)
	id := params["id"]
	id = strings.Trim(id, "{}")

	if id == "" {
		log.Errorf("Handler: DeleteTask: %v", fmt.Errorf("invalid id"))
		http.Error(w, "invalid id", http.StatusBadRequest)

		return
	}

	info := &e.EntityInfo{
		ID: id,
	}

	task, err := Delete(info)
	if err != nil {
		log.Infof("Handler: DeleteTask: %v", err)
		http.Error(w, err.Error(), http.StatusNotFound)

		return
	}

	taskStatus, ok := (*task).(e.TaskStatus)
	if !ok {
		err := fmt.Errorf("invalid data")
		log.Errorf("Handler: DeleteTask: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	data, err := json.Marshal(taskStatus)
	if err != nil {
		log.Errorf("Handler: DeleteTask: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	log.Tracef("Handler: GetTask")

	tasks, err := GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)

		return
	}

	tasksStatus := make([]e.TaskStatus, 0)
	for _, task := range tasks {
		taskStatus, ok := task.(e.TaskStatus)
		if !ok {
			err := fmt.Errorf("invalid data")
			log.Errorf("Handler: GetTask: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		if taskStatus.TaskInfo != nil {
			tasksStatus = append(tasksStatus, taskStatus)
		}
	}

	data, err := json.Marshal(tasksStatus)
	if err != nil {
		log.Errorf("Handler: GetTask: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
