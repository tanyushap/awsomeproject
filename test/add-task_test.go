package test

import (
	"awesomeProject/pkg/model"
	"awesomeProject/pkg/restcalls"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http/httptest"
	"testing"
)

func TestAddTaskHappyPath(t *testing.T) {

	model.AllTasks = []model.Todo{
		{Id: 1, TaskName: "Task 1", Description: "Task One", Status: "Pending", Creator: "Bob"},
		{Id: 2, TaskName: "Task 2", Description: "Task Two", Status: "Completed", Creator: "Bob"},
	}

	taskToAdd := model.Todo{Id: 3, TaskName: "Task 2", Description: "Task Two", Creator: "Bob"}

	newTaskJson, err := json.Marshal(taskToAdd)
	if err != nil {
		log.Printf("Failed to marshal json")
	}
	req := httptest.NewRequest("POST", "http://127.0.0.1:8081/task/add", bytes.NewBuffer(newTaskJson))
	writer := httptest.NewRecorder()

	restcalls.AddTask(writer, req)

	response := writer.Result()
	assert.NotNil(t, response)

	assert.Equal(t, "200 OK", response.Status)
	assert.Equal(t, 3, len(model.AllTasks))

}

func TestAddTaskWrongRequestType(t *testing.T) {

	model.AllTasks = []model.Todo{
		{Id: 1, TaskName: "Task 1", Description: "Task One", Status: "Pending", Creator: "Bob"},
		{Id: 2, TaskName: "Task 2", Description: "Task Two", Status: "Completed", Creator: "Bob"},
	}

	request := httptest.NewRequest("GET", "http://127.0.0.1:8081/task/add", nil)
	newRecorder := httptest.NewRecorder()

	restcalls.AddTask(newRecorder, request)

	response := newRecorder.Result()
	assert.NotNil(t, response)

	assert.Equal(t, "400 Bad Request", response.Status)
	assert.Equal(t, 2, len(model.AllTasks))

}

func TestAddTaskNilObjectWillReturnBadRequest(t *testing.T) {

	model.AllTasks = []model.Todo{
		{Id: 1, TaskName: "Task 1", Description: "Task One", Status: "Pending", Creator: "Bob"},
		{Id: 2, TaskName: "Task 2", Description: "Task Two", Status: "Completed", Creator: "Bob"},
	}

	request := httptest.NewRequest("POST", "http://127.0.0.1:8081/task/add", nil)
	newRecorder := httptest.NewRecorder()

	restcalls.AddTask(newRecorder, request)

	response := newRecorder.Result()
	assert.NotNil(t, response)

	assert.Equal(t, "400 Bad Request", response.Status)
	assert.Equal(t, 2, len(model.AllTasks))

}
