package restcalls

import (
	"awesomeProject/pkg/model"
	"encoding/json"
	"net/http"
)

func GetAllCompletedTasks(writer http.ResponseWriter, request *http.Request) {

	if request.Method == "GET" {
		getAllCompletedTasks(writer)
	} else {
		writer.WriteHeader(http.StatusBadRequest)
	}

}

func getAllCompletedTasks(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")

	var pendingTasks []model.Todo

	for _, task := range model.AllTasks {
		if task.Status == "Completed" {
			pendingTasks = append(pendingTasks, task)
		}
	}

	err := json.NewEncoder(writer).Encode(pendingTasks)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	writer.WriteHeader(http.StatusOK)
}
