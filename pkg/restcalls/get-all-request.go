package restcalls

import (
	"awesomeProject/pkg/model"
	"encoding/json"
	"net/http"
)

func GetAllTasks(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(writer).Encode(model.AllTasks)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)

}
