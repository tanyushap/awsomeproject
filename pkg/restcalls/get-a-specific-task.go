package restcalls

import (
	"awesomeProject/pkg/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func GetATaskById(writer http.ResponseWriter, request *http.Request) {

	var taskFound model.Todo
	var idWanted int64
	var err error
	requestParam := request.URL.Query()["id"]

	if requestParam != nil {
		idWanted, err = strconv.ParseInt(requestParam[0], 10, 64)
	} else {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("Looking for task with id %d", idWanted)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	isFound := false
	for _, task := range model.AllTasks {
		if task.Id == idWanted {
			isFound = true
			taskFound = task
			log.Printf("Task found : %s", taskFound)
			log.Print("Task found")
			err = json.NewEncoder(writer).Encode(taskFound)
			writer.WriteHeader(http.StatusOK)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusBadRequest)
				return
			}
		}

	}

	if !isFound {
		log.Print("Task not found")
		writer.WriteHeader(http.StatusNotFound)

	}

}
