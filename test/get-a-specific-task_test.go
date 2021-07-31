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

func TestGetASpecificTask_HappyPath(t *testing.T) {

	model.AllTasks = []model.Todo{
		{Id: 1, TaskName: "Task 1", Description: "Task One", Status: "Pending", Creator: "Bob"},
		{Id: 2, TaskName: "Task 2", Description: "Task Two", Status: "Completed", Creator: "Bob"},
	}

	req := httptest.NewRequest("GET", "http://127.0.0.1:8081/task?id=1", nil)
	writer := httptest.NewRecorder()

	restcalls.GetATaskById(writer, req)

	response := writer.Result()

	assert.NotNil(t, response)

	body, _ := io.ReadAll(response.Body)

	log.Printf("Response body: %s", body)

	assert.NotNil(t, body)

	var task model.Todo
	err := json.Unmarshal(body, &task)
	assert.Nil(t, err)

	assert.Equal(t, model.AllTasks[0], task)

}

func TestGetASpecificTaskReturns404ForNonExistentId_HappyPath(t *testing.T) {

	model.AllTasks = []model.Todo{
		{Id: 1, TaskName: "Task 1", Description: "Task One", Status: "Pending", Creator: "Bob"},
		{Id: 2, TaskName: "Task 2", Description: "Task Two", Status: "Completed", Creator: "Bob"},
	}

	req := httptest.NewRequest("GET", "http://127.0.0.1:8081/task?id=256", nil)
	writer := httptest.NewRecorder()

	restcalls.GetATaskById(writer, req)

	response := writer.Result()

	assert.NotNil(t, response)
	assert.Equal(t, "404 Not Found", response.Status)

}

func TestGetASpecificTaskReturns404ForNoId_HappyPath(t *testing.T) {

	model.AllTasks = []model.Todo{
		{Id: 1, TaskName: "Task 1", Description: "Task One", Status: "Pending", Creator: "Bob"},
		{Id: 2, TaskName: "Task 2", Description: "Task Two", Status: "Completed", Creator: "Bob"},
	}

	req := httptest.NewRequest("GET", "http://127.0.0.1:8081/task", nil)
	writer := httptest.NewRecorder()

	restcalls.GetATaskById(writer, req)

	response := writer.Result()

	assert.NotNil(t, response)
	assert.Equal(t, "400 Bad Request", response.Status)

}
