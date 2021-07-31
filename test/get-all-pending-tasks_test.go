package test

import (
	"awesomeProject/pkg/model"
	"awesomeProject/pkg/restcalls"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http/httptest"
	"testing"
)

func Test_GetAllPendingTasks_HappyPath(t *testing.T) {
	model.AllTasks = []model.Todo{
		{Id: 1, TaskName: "Task 1", Description: "Task One", Status: "Pending", Creator: "Bob"},
		{Id: 2, TaskName: "Task 2", Description: "Task Two", Status: "Completed", Creator: "Bob"},
	}

	req := httptest.NewRequest("GET", "http://127.0.0.1:8081/task/get-all-todo", nil)
	writer := httptest.NewRecorder()

	restcalls.GetAllPendingTasks(writer, req)

	response := writer.Result()

	assert.NotNil(t, response)

	body, _ := io.ReadAll(response.Body)

	log.Printf("Response body: %s", body)

	assert.NotNil(t, body)

	var tasks []model.Todo
	err := json.Unmarshal(body, &tasks)
	assert.Nil(t, err)

	assert.NotNil(t, tasks)
	assert.Equal(t, 1, len(tasks))
	assert.Equal(t, model.AllTasks[0], tasks[0])

}
