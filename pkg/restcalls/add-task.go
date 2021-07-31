package restcalls

import (
	"awesomeProject/pkg/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func AddTask(writer http.ResponseWriter, request *http.Request) {

	if request.Method == "POST" {
		log.Printf("Request type is POST")

		addNewTaskToTheTaskList(writer, request)
		return
	} else {

		writer.WriteHeader(http.StatusBadRequest)

		_, err := fmt.Fprintf(writer, "This Endpoint requires a POST Request!")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
	}

}

func addNewTaskToTheTaskList(writer http.ResponseWriter, request *http.Request) {

	var newTask model.Todo

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&newTask)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if &newTask != nil {
		newTask.Status = "Pending"
		model.AllTasks = append(model.AllTasks, newTask)
		writer.WriteHeader(http.StatusOK)
	}

}
