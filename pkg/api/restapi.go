package api

import (
	"awesomeProject/pkg/model"
	"awesomeProject/pkg/restcalls"
	"log"
	"net/http"
)

func HandleRestCalls() {

	model.AllTasks = []model.Todo{
		{Id: 1, TaskName: "Task 1", Description: "Task One", Status: "Pending", Creator: "Bob"},
		{Id: 2, TaskName: "Task 2", Description: "Task Two", Status: "Completed", Creator: "Bob"},
	}

	http.HandleFunc("/task/get-all", restcalls.GetAllTasks)
	http.HandleFunc("/task/get-all-todo", restcalls.GetAllPendingTasks)
	http.HandleFunc("/task/get-all-completed", restcalls.GetAllCompletedTasks)
	http.HandleFunc("/task", restcalls.GetATaskById)

	http.HandleFunc("/task/add", restcalls.AddTask)

	//http.HandleFunc("/task/complete", completeTask)
	//
	//http.HandleFunc("/task/delete", deleteTheTodoTask)

	log.Fatal(http.ListenAndServe(":8081", nil))

}
