package restcalls

import (
	"awesomeProject/pkg/model"
	"encoding/json"
	"net/http"
)

func GetAllPendingTasks(writer http.ResponseWriter, request *http.Request) {

	if request.Method == "GET" {
		getAllPendingTasks(writer)
	} else {
		writer.WriteHeader(http.StatusBadRequest)
	}

}

func getAllPendingTasks(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")

	var pendingTasks []model.Todo

	for _, task := range model.AllTasks {
		if task.Status == "Pending" {
			pendingTasks = append(pendingTasks, task)
		}
	}

	err := json.NewEncoder(writer).Encode(pendingTasks)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}
